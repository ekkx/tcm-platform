package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/service"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type Service interface {
	GetReservationByID(ctx context.Context, reservationID ulid.ULID) (*entity.Reservation, error)
	ListReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error)
}

type ServiceImpl struct {
	reservationRepo repository.Repository
	userService     service.Service
}

func New(reservationRepo repository.Repository, userService service.Service) Service {
	return &ServiceImpl{
		reservationRepo: reservationRepo,
		userService:     userService,
	}
}
