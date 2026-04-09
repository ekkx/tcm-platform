package main

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/gateway"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
	"github.com/ekkx/tcmrsv"
)

// --- Mock OfficialClient ---

type mockClient struct {
	loginErr   error
	reserveErr error
	// 呼び出し記録
	loginCalls   []*tcmrsv.LoginParams
	reserveCalls []*tcmrsv.ReserveParams
}

func (m *mockClient) Login(params *tcmrsv.LoginParams) error {
	m.loginCalls = append(m.loginCalls, params)
	return m.loginErr
}

func (m *mockClient) Reserve(params *tcmrsv.ReserveParams) error {
	m.reserveCalls = append(m.reserveCalls, params)
	return m.reserveErr
}

// --- Mock ReservationCommand ---

type statusUpdate struct {
	ID     ulid.ULID
	Status valueobject.ReservationStatus
}

type mockReservationCommand struct {
	mu      sync.Mutex
	updates []statusUpdate
	err     error
}

func (m *mockReservationCommand) UpdateReservationStatus(ctx context.Context, params *gateway.UpdateReservationStatusCommand) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.updates = append(m.updates, statusUpdate{
		ID:     params.ID,
		Status: params.Status,
	})
	return m.err
}

// --- Mock Notifier ---

type mockNotifier struct {
	mu   sync.Mutex
	msgs []NotifyMessage
}

func (m *mockNotifier) Notify(msg NotifyMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.msgs = append(m.msgs, msg)
}

// --- Helper ---

func ptr(s string) *string { return &s }

func makeGroup(officialID, password string, reservations ...*entity.Reservation) *Group {
	rsvViews := make([]*assemble.ReservationView, len(reservations))
	for i, rsv := range reservations {
		rsvViews[i] = &assemble.ReservationView{
			Reservation: *rsv,
			UserView:    assemble.UserView{},
		}
	}
	return &Group{
		MasterUser: entity.User{
			ID:                   ulid.New(),
			OfficialSiteID:       ptr(officialID),
			OfficialSitePassword: ptr(password),
		},
		Reservations: rsvViews,
	}
}

func makeReservation() *entity.Reservation {
	return &entity.Reservation{
		ID:         ulid.New(),
		CampusType: valueobject.CampusTypeIkebukuro,
		RoomID:     "101",
		Date:       ymd.New(2026, 4, 9),
		FromHour:   10,
		FromMinute: 0,
		ToHour:     11,
		ToMinute:   0,
		Status:     valueobject.ReservationStatusPending,
	}
}

// --- processGroup tests ---

func TestProcessGroup_AllSuccess(t *testing.T) {
	client := &mockClient{}
	factory := func() OfficialClient { return client }
	cmd := &mockReservationCommand{}

	rsv1 := makeReservation()
	rsv2 := makeReservation()
	group := makeGroup("user1", "pass1", rsv1, rsv2)

	result := processGroup(context.Background(), group, factory, cmd, &NoopNotifier{})

	// Login が1回呼ばれる
	if len(client.loginCalls) != 1 {
		t.Fatalf("expected 1 login call, got %d", len(client.loginCalls))
	}
	if client.loginCalls[0].UserID != "user1" {
		t.Errorf("expected login user_id 'user1', got '%s'", client.loginCalls[0].UserID)
	}

	// Reserve が2回呼ばれる
	if len(client.reserveCalls) != 2 {
		t.Fatalf("expected 2 reserve calls, got %d", len(client.reserveCalls))
	}

	// 両方 Success に更新される
	if len(cmd.updates) != 2 {
		t.Fatalf("expected 2 status updates, got %d", len(cmd.updates))
	}
	for _, u := range cmd.updates {
		if u.Status != valueobject.ReservationStatusSuccess {
			t.Errorf("expected status Success, got %d", u.Status)
		}
	}

	// GroupResult の内容
	if result.succeeded() != 2 {
		t.Errorf("expected 2 succeeded, got %d", result.succeeded())
	}
	if result.failed() != 0 {
		t.Errorf("expected 0 failed, got %d", result.failed())
	}
}

func TestProcessGroup_LoginFailure(t *testing.T) {
	client := &mockClient{loginErr: errors.New("auth failed")}
	factory := func() OfficialClient { return client }
	cmd := &mockReservationCommand{}

	rsv1 := makeReservation()
	rsv2 := makeReservation()
	group := makeGroup("user1", "pass1", rsv1, rsv2)

	result := processGroup(context.Background(), group, factory, cmd, &NoopNotifier{})

	// Reserve は呼ばれない
	if len(client.reserveCalls) != 0 {
		t.Errorf("expected 0 reserve calls, got %d", len(client.reserveCalls))
	}

	// 両方 Failed に更新される
	if len(cmd.updates) != 2 {
		t.Fatalf("expected 2 status updates, got %d", len(cmd.updates))
	}
	for _, u := range cmd.updates {
		if u.Status != valueobject.ReservationStatusFailed {
			t.Errorf("expected status Failed, got %d", u.Status)
		}
	}

	// GroupResult にログインエラーが記録される
	if result.LoginErr == nil {
		t.Error("expected LoginErr to be set")
	}
	if result.failed() != 2 {
		t.Errorf("expected 2 failed, got %d", result.failed())
	}
}

func TestProcessGroup_ReservePartialFailure(t *testing.T) {
	client := &partialFailClient{}
	factory := func() OfficialClient { return client }
	cmd := &mockReservationCommand{}

	rsv1 := makeReservation()
	rsv2 := makeReservation()
	group := makeGroup("user1", "pass1", rsv1, rsv2)

	result := processGroup(context.Background(), group, factory, cmd, &NoopNotifier{})

	// 2件の status 更新: 1件 Success, 1件 Failed
	if len(cmd.updates) != 2 {
		t.Fatalf("expected 2 status updates, got %d", len(cmd.updates))
	}

	successCount := 0
	failedCount := 0
	for _, u := range cmd.updates {
		switch u.Status {
		case valueobject.ReservationStatusSuccess:
			successCount++
		case valueobject.ReservationStatusFailed:
			failedCount++
		}
	}
	if successCount != 1 || failedCount != 1 {
		t.Errorf("expected 1 success and 1 failed, got %d success and %d failed", successCount, failedCount)
	}

	// GroupResult に1成功1失敗
	if result.succeeded() != 1 {
		t.Errorf("expected 1 succeeded, got %d", result.succeeded())
	}
	if result.failed() != 1 {
		t.Errorf("expected 1 failed, got %d", result.failed())
	}
}

// partialFailClient は1回目の Reserve は成功、2回目以降は失敗する mock
type partialFailClient struct {
	mu         sync.Mutex
	callCount  int
	loginCalls []*tcmrsv.LoginParams
}

func (c *partialFailClient) Login(params *tcmrsv.LoginParams) error {
	c.loginCalls = append(c.loginCalls, params)
	return nil
}

func (c *partialFailClient) Reserve(params *tcmrsv.ReserveParams) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.callCount++
	if c.callCount > 1 {
		return errors.New("reservation failed on official site")
	}
	return nil
}

func TestProcessGroup_MissingCredentials(t *testing.T) {
	client := &mockClient{}
	factory := func() OfficialClient { return client }
	cmd := &mockReservationCommand{}

	rsv1 := makeReservation()
	group := &Group{
		MasterUser: entity.User{
			ID:                   ulid.New(),
			OfficialSiteID:       nil, // credentials missing
			OfficialSitePassword: nil,
		},
		Reservations: []*assemble.ReservationView{
			{Reservation: *rsv1, UserView: assemble.UserView{}},
		},
	}

	result := processGroup(context.Background(), group, factory, cmd, &NoopNotifier{})

	// Login は呼ばれない
	if len(client.loginCalls) != 0 {
		t.Errorf("expected 0 login calls, got %d", len(client.loginCalls))
	}

	// Failed に更新される
	if len(cmd.updates) != 1 {
		t.Fatalf("expected 1 status update, got %d", len(cmd.updates))
	}
	if cmd.updates[0].Status != valueobject.ReservationStatusFailed {
		t.Errorf("expected status Failed, got %d", cmd.updates[0].Status)
	}

	// GroupResult に認証情報未設定が記録される
	if !result.MissingCreds {
		t.Error("expected MissingCreds to be true")
	}
	if result.failed() != 1 {
		t.Errorf("expected 1 failed, got %d", result.failed())
	}
}

// --- buildSummaryMessage tests ---

func TestBuildSummaryMessage_AllSuccess(t *testing.T) {
	results := []GroupResult{
		{
			DisplayID: "user1",
			Items: []ReservationItemResult{
				{RoomID: "101", Date: "2026-04-11", FromHour: 10, FromMinute: 0, ToHour: 11, ToMinute: 0, Success: true},
				{RoomID: "102", Date: "2026-04-11", FromHour: 14, FromMinute: 0, ToHour: 15, ToMinute: 0, Success: true},
			},
		},
	}

	msg := buildSummaryMessage("2026-04-11", results)

	if msg.Level != NotifyLevelInfo {
		t.Errorf("expected Info level, got %d", msg.Level)
	}
	if !strings.Contains(msg.Title, "成功 2") {
		t.Errorf("expected title to contain '成功 2', got %q", msg.Title)
	}
	if !strings.Contains(msg.Title, "失敗 0") {
		t.Errorf("expected title to contain '失敗 0', got %q", msg.Title)
	}

	// user1 のフィールドが含まれる
	userField := findField(msg.Fields, "👤 user1")
	if userField == nil {
		t.Fatal("expected field for '👤 user1'")
	}
	if !strings.Contains(userField.Value, "✅") {
		t.Errorf("expected field value to contain '✅', got %q", userField.Value)
	}
	if strings.Contains(userField.Value, "❌") {
		t.Errorf("unexpected '❌' in field value: %q", userField.Value)
	}
}

func TestBuildSummaryMessage_WithFailures(t *testing.T) {
	results := []GroupResult{
		{
			DisplayID: "user1",
			Items: []ReservationItemResult{
				{RoomID: "101", FromHour: 10, ToHour: 11, Success: true},
				{RoomID: "102", FromHour: 14, ToHour: 15, Success: false, Err: errors.New("room not available")},
			},
		},
	}

	msg := buildSummaryMessage("2026-04-11", results)

	if msg.Level != NotifyLevelInfo {
		t.Errorf("expected Error level when there are failures, got %d", msg.Level)
	}
	if !strings.Contains(msg.Title, "成功 1") {
		t.Errorf("expected title to contain '成功 1', got %q", msg.Title)
	}
	if !strings.Contains(msg.Title, "失敗 1") {
		t.Errorf("expected title to contain '失敗 1', got %q", msg.Title)
	}

	userField := findField(msg.Fields, "👤 user1")
	if userField == nil {
		t.Fatal("expected field for '👤 user1'")
	}
	if !strings.Contains(userField.Value, "✅") {
		t.Errorf("expected '✅' in field value, got %q", userField.Value)
	}
	if !strings.Contains(userField.Value, "❌") {
		t.Errorf("expected '❌' in field value, got %q", userField.Value)
	}
	if !strings.Contains(userField.Value, "(room not available)") {
		t.Errorf("expected error message in field value, got %q", userField.Value)
	}
}

func TestBuildSummaryMessage_LoginFailure(t *testing.T) {
	results := []GroupResult{
		{
			DisplayID: "user1",
			LoginErr:  errors.New("auth failed"),
			Items: []ReservationItemResult{
				{RoomID: "101", Success: false},
			},
		},
	}

	msg := buildSummaryMessage("2026-04-11", results)

	if msg.Level != NotifyLevelInfo {
		t.Errorf("expected Error level, got %d", msg.Level)
	}

	userField := findField(msg.Fields, "👤 user1")
	if userField == nil {
		t.Fatal("expected field for '👤 user1'")
	}
	if !strings.Contains(userField.Value, "ログイン失敗") {
		t.Errorf("expected 'ログイン失敗' in field value, got %q", userField.Value)
	}
	if !strings.Contains(userField.Value, "auth failed") {
		t.Errorf("expected error message in field value, got %q", userField.Value)
	}
}

func TestBuildSummaryMessage_MissingCredentials(t *testing.T) {
	results := []GroupResult{
		{
			DisplayID:    "user1",
			MissingCreds: true,
			Items: []ReservationItemResult{
				{RoomID: "101", Success: false},
			},
		},
	}

	msg := buildSummaryMessage("2026-04-11", results)

	if msg.Level != NotifyLevelInfo {
		t.Errorf("expected Error level, got %d", msg.Level)
	}

	userField := findField(msg.Fields, "👤 user1")
	if userField == nil {
		t.Fatal("expected field for '👤 user1'")
	}
	if !strings.Contains(userField.Value, "認証情報が未設定") {
		t.Errorf("expected '認証情報が未設定' in field value, got %q", userField.Value)
	}
}

func TestBuildSummaryMessage_MultipleUsers(t *testing.T) {
	results := []GroupResult{
		{
			DisplayID: "user1",
			Items: []ReservationItemResult{
				{RoomID: "101", Success: true},
			},
		},
		{
			DisplayID: "user2",
			LoginErr:  errors.New("auth failed"),
			Items: []ReservationItemResult{
				{RoomID: "201", Success: false},
			},
		},
	}

	msg := buildSummaryMessage("2026-04-11", results)

	if msg.Level != NotifyLevelInfo {
		t.Errorf("expected Error level, got %d", msg.Level)
	}
	if !strings.Contains(msg.Title, "成功 1") {
		t.Errorf("expected '成功 1' in title, got %q", msg.Title)
	}
	if !strings.Contains(msg.Title, "失敗 1") {
		t.Errorf("expected '失敗 1' in title, got %q", msg.Title)
	}
	if findField(msg.Fields, "👤 user1") == nil {
		t.Error("expected field for '👤 user1'")
	}
	if findField(msg.Fields, "👤 user2") == nil {
		t.Error("expected field for '👤 user2'")
	}
}

// --- helper ---

func findField(fields []NotifyField, name string) *NotifyField {
	for i := range fields {
		if fields[i].Name == name {
			return &fields[i]
		}
	}
	return nil
}
