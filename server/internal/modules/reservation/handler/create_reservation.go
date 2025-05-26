package handler

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/core/vo"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/room"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) CreateReservation(
	ctx context.Context,
	req *reservation_v1.CreateReservationRequest,
) (*reservation_v1.CreateReservationReply, error) {
	var date time.Time
	if req.Reservation.Date != nil {
		date = req.Reservation.Date.AsTime()
	}

	var pianoTypes []vo.PianoType
	if req.Reservation.PianoTypes != nil {
		pianoTypes = make([]vo.PianoType, len(req.Reservation.PianoTypes))
		for i, pianoType := range req.Reservation.PianoTypes {
			pianoTypes[i] = vo.PianoType(pianoType.Number())
		}
	}

	rsv, err := h.Usecase.CreateReservation(ctx, &usecase.CreateReservationInput{
		// UserID: ctxhelper.GetGRPCRequestUser(ctx),
		Campus:       vo.Campus(req.Reservation.Campus.String()),
		Date:         &date,
		FromHour:     req.Reservation.FromHour,
		FromMinute:   req.Reservation.FromMinute,
		ToHour:       req.Reservation.ToHour,
		ToMinute:     req.Reservation.ToMinute,
		IsAutoSelect: req.Reservation.IsAutoSelect,
		RoomId:       req.Reservation.RoomId,
		BookerName:   req.Reservation.BookerName,
		PianoNumbers: req.Reservation.PianoNumbers,
		PianoTypes:   pianoTypes,
		Floors:       req.Reservation.Floors,
		IsBasement:   req.Reservation.IsBasement,
	})
	if err != nil {
		return nil, err
	}

	var campus room_v1.Campus
	switch rsv.Reservation.Campus {
	case vo.CampusNakameguro:
		campus = room_v1.Campus_NAKAMEGURO
	case vo.CampusIkebukuro:
		campus = room_v1.Campus_IKEBUKURO
	default:
		campus = room_v1.Campus_CAMPUS_UNSPECIFIED
	}

	return &reservation_v1.CreateReservationReply{
		Reservation: &reservation_v1.Reservation{
			Id:         rsv.Reservation.ID,
			ExternalId: rsv.Reservation.ExternalID,
			Campus:     campus,
			Date:       timestamppb.New(rsv.Reservation.Date),
			RoomId:     rsv.Reservation.RoomID,
			FromHour:   rsv.Reservation.FromHour,
			FromMinute: rsv.Reservation.FromMinute,
			ToHour:     rsv.Reservation.ToHour,
			ToMinute:   rsv.Reservation.ToMinute,
			BookerName: rsv.Reservation.BookerName,
			CreatedAt:  timestamppb.New(rsv.Reservation.CreatedAt),
		},
	}, nil
}
