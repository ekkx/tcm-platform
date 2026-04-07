package adapter

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/app/gateway"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
)

type CommandAdapter struct {
	reservationRepo repository.Repository
}

func NewCommandAdapter(reservationRepo repository.Repository) gateway.ReservationCommand {
	return &CommandAdapter{
		reservationRepo: reservationRepo,
	}
}

func (a *CommandAdapter) UpdateReservationStatus(ctx context.Context, params *gateway.UpdateReservationStatusCommand) error {
	return a.reservationRepo.UpdateReservationStatus(ctx, &repository.UpdateReservationStatusParams{
		ID:             params.ID,
		Status:         params.Status,
		OfficialSiteID: params.OfficialSiteID,
	})
}
