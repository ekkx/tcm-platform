package input_test

import (
	"context"
	"testing"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetUserReservations_FromProto(t *testing.T) {
	testUserID := "test-user-123"

	tests := []struct {
		name     string
		req      *reservation.GetUserReservationsRequest
		userID   string
		expected func(*testing.T, *input.GetUserReservations)
	}{
		{
			name: "FromDateがnilの場合",
			req: &reservation.GetUserReservationsRequest{
				FromDate: nil,
			},
			userID: testUserID,
			expected: func(t *testing.T, result *input.GetUserReservations) {
				assert.Equal(t, testUserID, result.UserID)
				assert.True(t, result.FromDate.IsZero(), "FromDate should be zero value when nil")
			},
		},
		{
			name: "FromDateが指定されている場合",
			req: &reservation.GetUserReservationsRequest{
				FromDate: timestamppb.New(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			userID: testUserID,
			expected: func(t *testing.T, result *input.GetUserReservations) {
				assert.Equal(t, testUserID, result.UserID)
				assert.False(t, result.FromDate.IsZero(), "FromDate should not be zero value")
				assert.Equal(t, 2024, result.FromDate.Year())
				assert.Equal(t, time.January, result.FromDate.Month())
				assert.Equal(t, 1, result.FromDate.Day())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// コンテキストにアクター情報を設定
			ctx := context.Background()
			ctx = ctxhelper.SetActor(ctx, actor.Actor{
				ID:   tt.userID,
				Role: actor.RoleUser,
			})

			// テスト実行
			dto := input.NewGetUserReservations()
			result := dto.FromProto(ctx, tt.req)

			// 検証
			tt.expected(t, result)
		})
	}
}

func TestGetUserReservations_Validate(t *testing.T) {
	tests := []struct {
		name      string
		input     *input.GetUserReservations
		wantError bool
		errorMsg  string
	}{
		{
			name: "FromDateがゼロ値の場合は正常",
			input: &input.GetUserReservations{
				UserID:   "test-user",
				FromDate: time.Time{},
			},
			wantError: false,
		},
		{
			name: "FromDateが将来の日付の場合は正常",
			input: &input.GetUserReservations{
				UserID:   "test-user",
				FromDate: time.Now().Add(24 * time.Hour),
			},
			wantError: false,
		},
		{
			name: "FromDateが過去の日付の場合はエラー",
			input: &input.GetUserReservations{
				UserID:   "test-user",
				FromDate: time.Now().Add(-24 * time.Hour),
			},
			wantError: true,
			errorMsg:  "date_must_be_today_or_future",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()

			if tt.wantError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
