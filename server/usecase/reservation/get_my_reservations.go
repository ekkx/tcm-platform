package reservation

import (
	"context"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/domain"
	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
)

type GetMyReservationsInput struct {
	UserID   string
	Password string
}

type GetMyReservationsOutput struct {
	Reservations []domain.Reservation
}

func (uc *ReservationUsecaseImpl) GetMyReservations(ctx context.Context, input *GetMyReservationsInput) (*GetMyReservationsOutput, error) {
	if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
		UserID:   input.UserID,
		Password: input.Password,
	}); err != nil {
		return nil, apperrors.ErrUnauthorized
	}

	// TODO: 予約データはすでにデータベースに保存されているので、DBから取得するようにする
	reservations, err := uc.tcmClient.GetMyReservations()
	if err != nil {
		return nil, err
	}

	rooms := uc.tcmClient.GetRooms()

	roomNameToID := make(map[string]string, len(rooms))
	for _, room := range rooms {
		roomNameToID[room.Name] = room.ID
	}

	var domainReservations []domain.Reservation
	for _, rsv := range reservations {
		domainReservations = append(domainReservations, domain.Reservation{
			ID:     rsv.ID,
			RoomID: roomNameToID[rsv.RoomName],
			// TODO: map other fields
		})
	}

	return &GetMyReservationsOutput{
		Reservations: domainReservations,
	}, nil
}
