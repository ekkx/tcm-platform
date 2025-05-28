package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
)

func (uc *Usecase) GetRooms(ctx context.Context) *output.GetRooms {
	// 全て取ってきたいので検索クエリは空のまま
	rooms := uc.roomRepo.SearchRooms(ctx, repository.SearchRoomsArgs{})
	return output.NewGetRooms(rooms)
}
