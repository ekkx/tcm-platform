package reservation

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/adapter/db/mapper"
	"github.com/ekkx/tcmrsv-web/server/domain"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
)

type CreateReservationInput struct {
	UserID   string
	Password string

	Campus       domain.Campus
	Date         time.Time
	FromHour     int
	FromMinute   int
	ToHour       int
	ToMinute     int
	IsAutoSelect bool

	RoomID       *string
	BookerName   *string
	PianoNumbers *[]int
	PianoTypes   *[]domain.PianoType
	Floors       *[]int
	IsBasement   *bool
}

type CreateReservationOutput struct {
	Reservation domain.Reservation
}

func (uc ReservationUsecaseImpl) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
		UserID:   input.UserID,
		Password: input.Password,
	}); err != nil {
		return nil, apperrors.ErrUnauthorized
	}

	var roomID string
	if input.IsAutoSelect {
		// TODO: 条件に合う練習室を検索してランダムに一つ選ぶ（すでに存在する予約の部屋と時間が一致しないこと）
	} else {
		if input.RoomID == nil {
			return nil, apperrors.ErrRoomIDRequired
		}

		// 予約の部屋が存在するか確認する
		var room *tcmrsv.Room
		for _, r := range uc.tcmClient.GetRooms() {
			if r.ID == *input.RoomID {
				room = &r
				break
			}
		}

		if room == nil {
			return nil, apperrors.ErrRoomNotFound
		}

		roomID = *input.RoomID
	}

	// すでに存在する予約の部屋と時間が一致しないことを確認する
	hasConflict, err := uc.querier.CheckReservationConflict(ctx, db.CheckReservationConflictParams{
		RoomID: roomID,
		Date:   time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, utils.JST()),
		FromHour: func() *int32 {
			h := int32(input.FromHour)
			return &h
		}(),
		FromMinute: func() *int32 {
			m := int32(input.FromMinute)
			return &m
		}(),
		ToHour: func() *int32 {
			h := int32(input.ToHour)
			return &h
		}(),
		ToMinute: func() *int32 {
			m := int32(input.ToMinute)
			return &m
		}(),
	})
	if err != nil {
		return nil, err
	}

	if hasConflict {
		return nil, apperrors.ErrReservationConflict
	}

	rsv, err := uc.querier.CreateReservation(ctx, db.CreateReservationParams{
		UserID:     input.UserID,
		Campus:     db.Campus(input.Campus),
		RoomID:     roomID,
		Date:       time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, utils.JST()),
		FromHour:   int32(input.FromHour),
		FromMinute: int32(input.FromMinute),
		ToHour:     int32(input.ToHour),
		ToMinute:   int32(input.ToMinute),
		BookerName: input.BookerName,
	})
	if err != nil {
		return nil, err
	}

	return &CreateReservationOutput{
		Reservation: mapper.ToReservation(rsv),
	}, nil
}
