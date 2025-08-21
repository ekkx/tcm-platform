package assemble

import (
	"context"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/shared/gateway"
	"github.com/ekkx/tcmrsv-web/internal/shared/mapper"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type ReservationView struct {
	Reservation entity.Reservation
	UserView    UserView
	Room        entity.Room
}

type ReservationAssembler interface {
	Build(ctx context.Context, reservationID ulid.ULID) (*ReservationView, error)
	BuildList(ctx context.Context, reservationIDs []ulid.ULID) ([]*ReservationView, error)
}

type ReservationAssemblerImpl struct {
	reservationQuery gateway.ReservationQuery
	userAsm          UserAssembler
}

func NewReservationAssembler(reservationQuery gateway.ReservationQuery, userAsm UserAssembler) ReservationAssembler {
	return &ReservationAssemblerImpl{
		reservationQuery: reservationQuery,
		userAsm:          userAsm,
	}
}

func (asm *ReservationAssemblerImpl) Build(ctx context.Context, reservationID ulid.ULID) (*ReservationView, error) {
	v, err := asm.BuildList(ctx, []ulid.ULID{reservationID})
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, nil
	}
	return v[0], nil
}

func (asm *ReservationAssemblerImpl) BuildList(ctx context.Context, reservationIDs []ulid.ULID) ([]*ReservationView, error) {
	if len(reservationIDs) == 0 {
		return []*ReservationView{}, nil
	}

	entRsvs, err := asm.reservationQuery.ListReservationsByIDs(ctx, reservationIDs)
	if err != nil {
		return nil, err
	}

	if len(entRsvs) == 0 {
		return []*ReservationView{}, nil
	}

	// ユーザのバルク取得 & マップ化
	userIDSet := make(map[ulid.ULID]struct{}, len(entRsvs))
	roomIDSet := make(map[string]struct{}, len(entRsvs))
	for _, r := range entRsvs {
		userIDSet[r.UserID] = struct{}{}
		roomIDSet[r.RoomID] = struct{}{}
	}

	userIDs := make([]ulid.ULID, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	userViews, err := asm.userAsm.BuildList(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	userMap := make(map[ulid.ULID]*UserView, len(userViews))
	for _, uv := range userViews {
		if uv != nil {
			userMap[uv.User.ID] = uv
		}
	}

	// --- ルームのバルク取得 & マップ化 ---
	roomIDs := make([]string, 0, len(roomIDSet))
	for id := range roomIDSet {
		roomIDs = append(roomIDs, id)
	}

	rooms := tcmrsv.New().GetRooms()

	roomMap := make(map[string]*entity.Room, len(rooms))
	for _, rm := range rooms {
		roomMap[rm.ID] = mapper.ToRoom(&rm)
	}

	index := make(map[ulid.ULID]*entity.Reservation, len(entRsvs))
	for _, r := range entRsvs {
		index[r.ID] = r
	}

	out := make([]*ReservationView, 0, len(reservationIDs))
	for _, rid := range reservationIDs {
		if r := index[rid]; r != nil {
			out = append(out, &ReservationView{
				Reservation: *r,
				UserView:    *userMap[r.UserID],
				Room:        *roomMap[r.RoomID],
			})
		}
	}

	return out, nil
}
