package handler_test

import (
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
	"github.com/stretchr/testify/require"
)

// TestHandler_DTOConversion tests DTO conversion logic in handler
func TestHandler_DTOConversion(t *testing.T) {
	t.Run("output DTO のピアノ種別変換確認", func(t *testing.T) {

		// GetRoomsメソッドが内部でrepositoryを呼び出すため、
		// 実際のテストでは、repositoryをモックする必要がある
		// ここでは、DTO変換のテストのため、output.GetRoomsを直接テストする

		rooms := output.NewGetRooms([]entity.Room{
			{
				ID:          "room-1",
				CampusType:  enum.CampusTypeIkebukuro,
				Name:        "テスト部屋",
				PianoType:   enum.PianoTypeGrand,
				PianoNumber: 1,
			},
			{
				ID:          "room-2",
				CampusType:  enum.CampusTypeNakameguro,
				Name:        "ピアノなし部屋",
				PianoType:   enum.PianoTypeNone,
				PianoNumber: 0,
			},
		})

		reply := rooms.ToProto()
		require.NotNil(t, reply)
		require.Len(t, reply.Rooms, 2)

		// グランドピアノの変換確認
		require.Equal(t, room.PianoType(enum.PianoTypeGrand), reply.Rooms[0].PianoType)
		require.Equal(t, int32(1), reply.Rooms[0].PianoNumber)

		// ピアノなしの変換確認
		require.Equal(t, room.PianoType(enum.PianoTypeNone), reply.Rooms[1].PianoType)
		require.Equal(t, int32(0), reply.Rooms[1].PianoNumber)
	})

	t.Run("output DTO のキャンパス種別変換確認", func(t *testing.T) {
		rooms := output.NewGetRooms([]entity.Room{
			{
				ID:          "room-1",
				CampusType:  enum.CampusTypeIkebukuro,
				Name:        "池袋キャンパス部屋",
				PianoType:   enum.PianoTypeNone,
				PianoNumber: 0,
			},
			{
				ID:          "room-2",
				CampusType:  enum.CampusTypeNakameguro,
				Name:        "中目黒キャンパス部屋",
				PianoType:   enum.PianoTypeNone,
				PianoNumber: 0,
			},
		})

		reply := rooms.ToProto()
		require.NotNil(t, reply)
		require.Len(t, reply.Rooms, 2)

		// 池袋キャンパスの変換確認
		require.Equal(t, room.CampusType_IKEBUKURO, reply.Rooms[0].CampusType)

		// 中目黒キャンパスの変換確認
		require.Equal(t, room.CampusType_NAKAMEGURO, reply.Rooms[1].CampusType)
	})
}
