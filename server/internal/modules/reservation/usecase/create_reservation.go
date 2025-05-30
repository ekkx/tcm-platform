package usecase

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	rsv_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/apperrors"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
)

func isTimeConflict(rsv *entity.Reservation, params *input.CreateReservation) bool {
	return !(params.ToHour < rsv.FromHour ||
		(params.ToHour == rsv.FromHour && params.ToMinute <= rsv.FromMinute) ||
		(params.FromHour > rsv.ToHour ||
			(params.FromHour == rsv.ToHour && params.FromMinute >= rsv.ToMinute)))
}

type TimeSlot struct {
	FromHour   int32
	FromMinute int32
	ToHour     int32
	ToMinute   int32
}

func splitTimeSlot(params *input.CreateReservation) []TimeSlot {
	var slots []TimeSlot
	h := params.FromHour
	m := params.FromMinute

	for {
		nextH := h
		nextM := m + 30
		if nextM >= 60 {
			nextH++
			nextM = 0
		}
		if nextH > params.ToHour || (nextH == params.ToHour && nextM > params.ToMinute) {
			break
		}
		slots = append(slots, TimeSlot{
			FromHour:   h,
			FromMinute: m,
			ToHour:     nextH,
			ToMinute:   nextM,
		})
		h, m = nextH, nextM
	}
	return slots
}

func isSlotConflict(rsv *entity.Reservation, slot TimeSlot) bool {
	return !(slot.ToHour < rsv.FromHour ||
		(slot.ToHour == rsv.FromHour && slot.ToMinute <= rsv.FromMinute) ||
		(slot.FromHour > rsv.ToHour ||
			(slot.FromHour == rsv.ToHour && slot.FromMinute >= rsv.ToMinute)))
}

func (u *Usecase) CreateReservation(ctx context.Context, params *input.CreateReservation) (*output.CreateReservation, error) {
	var roomID string
	var date time.Time
	if params.Date != nil {
		date = time.Date(params.Date.Year(), params.Date.Month(), params.Date.Day(), 0, 0, 0, 0, utils.JST())
	}

	if params.IsAutoSelect {
		if params.Date == nil {
			return nil, apperrors.ErrReservationDateRequired
		}

		// 自動選択の場合、予約可能な練習室を取得
		rooms := u.roomRepo.SearchRooms(ctx, &room_repo.SearchRoomsArgs{
			PianoNumbers: params.PianoNumbers,
			PianoTypes:   params.PianoTypes,
			Floors:       params.Floors,
			IsBasement:   params.IsBasement,
			CampusTypes:  []enum.CampusType{params.CampusType},
		})
		if len(rooms) == 0 {
			return nil, apperrors.ErrRoomNotFound
		}

		// 指定された日付時間に存在する予約を取得
		rsvsByDate, err := u.rsvRepo.GetReservationsByDate(ctx, &rsv_repo.GetReservationsByDate{
			Date: date,
		})
		if err != nil {
			return nil, err
		}

		// 候補練習室の中で全時間帯が空いている部屋を優先
		candidateRoomID := ""
	ROOM_LOOP:
		for _, room := range rooms {
			for _, rsv := range rsvsByDate {
				if rsv.RoomID != room.ID {
					continue
				}
				if isTimeConflict(&rsv, params) {
					continue ROOM_LOOP // この練習室は時間帯が被っているのでスキップ
				}
			}
			candidateRoomID = room.ID
			break // 全ての時間帯が空いている練習室を見つけたのでループを抜ける
		}

		// 候補練習室が見つかった場合はそれを使用
		if candidateRoomID != "" {
			roomID = candidateRoomID
		} else {
			// 時間ごとに空いてる部屋を選ぶ（分散予約）
			var reservedList []entity.Reservation

			timeSlots := splitTimeSlot(params)
			for _, slot := range timeSlots {
				roomFound := false
				for _, room := range rooms {
					conflict := false
					for _, rsv := range rsvsByDate {
						if rsv.RoomID != room.ID {
							continue
						}
						if isSlotConflict(&rsv, slot) {
							conflict = true
							break // この練習室は時間帯が被っているのでスキップ
						}
					}
					if !conflict {
						// この練習室は空いているので予約可能
						r, err := u.rsvRepo.CreateReservation(ctx, &rsv_repo.CreateReservationArgs{
							UserID:     params.UserID,
							CampusType: params.CampusType,
							RoomID:     room.ID,
							Date:       date,
							FromHour:   slot.FromHour,
							FromMinute: slot.FromMinute,
							ToHour:     slot.ToHour,
							ToMinute:   slot.ToMinute,
							BookerName: params.BookerName,
						})
						if err != nil {
							return nil, err
						}
						reservedList = append(reservedList, r)
						roomFound = true
						break // この時間帯はこの練習室で予約できたので次の時間帯へ
					}
				}
				if !roomFound {
					// この時間帯に予約可能な練習室が見つからなかった場合はエラー
					return nil, apperrors.ErrNoAvailableRoom
				}
			}

			return output.NewCreateReservation(reservedList), nil
		}
	} else {
		// 手動選択の場合、指定された練習室IDを使用
		if params.RoomID == nil {
			return nil, apperrors.ErrRoomIDRequired
		}

		// 練習室の存在チェック
		rooms := u.roomRepo.SearchRooms(ctx, &room_repo.SearchRoomsArgs{
			ID: params.RoomID,
		})

		if len(rooms) == 0 {
			return nil, apperrors.ErrRoomNotFound
		}

		roomID = *params.RoomID
	}

	// 予約時間と練習室が被っていないか確認
	hasConflict, err := u.rsvRepo.CheckReservationConflict(ctx, &rsv_repo.CheckReservationConflictArgs{
		RoomID:     roomID,
		Date:       date,
		FromHour:   params.FromHour,
		FromMinute: params.FromMinute,
		ToHour:     params.ToHour,
		ToMinute:   params.ToMinute,
	})
	if err != nil {
		return nil, err
	}

	if hasConflict {
		return nil, apperrors.ErrReservationConflict
	}

	rsv, err := u.rsvRepo.CreateReservation(ctx, &rsv_repo.CreateReservationArgs{
		UserID:     params.UserID,
		CampusType: params.CampusType,
		RoomID:     roomID,
		Date:       date,
		FromHour:   params.FromHour,
		FromMinute: params.FromMinute,
		ToHour:     params.ToHour,
		ToMinute:   params.ToMinute,
		BookerName: params.BookerName,
	})
	if err != nil {
		return nil, err
	}

	return output.NewCreateReservation([]entity.Reservation{rsv}), nil
}
