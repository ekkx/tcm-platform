package main

import (
	"context"
	"errors"
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
			ID:                  ulid.New(),
			OfficialSiteID:      ptr(officialID),
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

// --- Tests ---

func TestProcessGroup_AllSuccess(t *testing.T) {
	client := &mockClient{}
	factory := func() OfficialClient { return client }
	cmd := &mockReservationCommand{}

	rsv1 := makeReservation()
	rsv2 := makeReservation()
	group := makeGroup("user1", "pass1", rsv1, rsv2)

	processGroup(context.Background(), group, factory, cmd)

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
}

func TestProcessGroup_LoginFailure(t *testing.T) {
	client := &mockClient{loginErr: errors.New("auth failed")}
	factory := func() OfficialClient { return client }
	cmd := &mockReservationCommand{}

	rsv1 := makeReservation()
	rsv2 := makeReservation()
	group := makeGroup("user1", "pass1", rsv1, rsv2)

	processGroup(context.Background(), group, factory, cmd)

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
}

func TestProcessGroup_ReservePartialFailure(t *testing.T) {
	client := &partialFailClient{}
	factory := func() OfficialClient { return client }
	cmd := &mockReservationCommand{}

	rsv1 := makeReservation()
	rsv2 := makeReservation()
	group := makeGroup("user1", "pass1", rsv1, rsv2)

	processGroup(context.Background(), group, factory, cmd)

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
			ID:                  ulid.New(),
			OfficialSiteID:      nil, // credentials missing
			OfficialSitePassword: nil,
		},
		Reservations: []*assemble.ReservationView{
			{Reservation: *rsv1, UserView: assemble.UserView{}},
		},
	}

	processGroup(context.Background(), group, factory, cmd)

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
}
