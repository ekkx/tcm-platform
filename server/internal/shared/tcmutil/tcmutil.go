package tcmutil

import (
	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/core/enum"
)

func ConvertRoomPianoType(pianoType tcmrsv.RoomPianoType) enum.PianoType {
	switch pianoType {
	case tcmrsv.RoomPianoTypeGrand:
		return enum.PianoTypeGrand
	case tcmrsv.RoomPianoTypeUpright:
		return enum.PianoTypeUpright
	case tcmrsv.RoomPianoTypeUnknown:
		return enum.PianoTypeUnknown
	default:
		return enum.PianoTypeNone
	}
}

func ConvertCampusType(campusType tcmrsv.Campus) enum.CampusType {
	switch campusType {
	case tcmrsv.CampusIkebukuro:
		return enum.CampusTypeIkebukuro
	case tcmrsv.CampusNakameguro:
		return enum.CampusTypeNakameguro
	case tcmrsv.CampusUnknown:
		return enum.CampusTypeUnknown
	default:
		return enum.CampusTypeUnknown
	}
}
