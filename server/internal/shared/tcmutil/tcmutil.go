package tcmutil

import (
	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
)

func ToTCMRoomPianoType(pianoType enum.PianoType) tcmrsv.RoomPianoType {
	switch pianoType {
	case enum.PianoTypeGrand:
		return tcmrsv.RoomPianoTypeGrand
	case enum.PianoTypeUpright:
		return tcmrsv.RoomPianoTypeUpright
	case enum.PianoTypeNone:
		return tcmrsv.RoomPianoTypeNone
	default:
		return tcmrsv.RoomPianoTypeUnknown
	}
}

func ToTCMCampusType(campusType enum.CampusType) tcmrsv.Campus {
	switch campusType {
	case enum.CampusTypeIkebukuro:
		return tcmrsv.CampusIkebukuro
	case enum.CampusTypeNakameguro:
		return tcmrsv.CampusNakameguro
	default:
		return tcmrsv.CampusUnknown
	}
}

func ToDomainPianoType(pianoType tcmrsv.RoomPianoType) enum.PianoType {
	switch pianoType {
	case tcmrsv.RoomPianoTypeGrand:
		return enum.PianoTypeGrand
	case tcmrsv.RoomPianoTypeUpright:
		return enum.PianoTypeUpright
	case tcmrsv.RoomPianoTypeNone:
		return enum.PianoTypeNone
	default:
		return enum.PianoTypeUnknown
	}
}

func ToDomainCampusType(campusType tcmrsv.Campus) enum.CampusType {
	switch campusType {
	case tcmrsv.CampusIkebukuro:
		return enum.CampusTypeIkebukuro
	case tcmrsv.CampusNakameguro:
		return enum.CampusTypeNakameguro
	default:
		return enum.CampusTypeUnknown
	}
}
