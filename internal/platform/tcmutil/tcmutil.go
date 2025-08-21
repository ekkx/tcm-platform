package tcmutil

import (
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcmrsv"
)

func ToTCMRoomPianoType(pianoType valueobject.PianoType) tcmrsv.RoomPianoType {
	switch pianoType {
	case valueobject.PianoTypeGrand:
		return tcmrsv.RoomPianoTypeGrand
	case valueobject.PianoTypeUpright:
		return tcmrsv.RoomPianoTypeUpright
	case valueobject.PianoTypeNone:
		return tcmrsv.RoomPianoTypeNone
	default:
		return tcmrsv.RoomPianoTypeUnknown
	}
}

func ToTCMCampusType(campusType valueobject.CampusType) tcmrsv.Campus {
	switch campusType {
	case valueobject.CampusTypeIkebukuro:
		return tcmrsv.CampusIkebukuro
	case valueobject.CampusTypeNakameguro:
		return tcmrsv.CampusNakameguro
	default:
		return tcmrsv.CampusUnknown
	}
}

func ToDomainPianoType(pianoType tcmrsv.RoomPianoType) valueobject.PianoType {
	switch pianoType {
	case tcmrsv.RoomPianoTypeGrand:
		return valueobject.PianoTypeGrand
	case tcmrsv.RoomPianoTypeUpright:
		return valueobject.PianoTypeUpright
	case tcmrsv.RoomPianoTypeNone:
		return valueobject.PianoTypeNone
	default:
		return valueobject.PianoTypeUnknown
	}
}

func ToDomainCampusType(campusType tcmrsv.Campus) valueobject.CampusType {
	switch campusType {
	case tcmrsv.CampusIkebukuro:
		return valueobject.CampusTypeIkebukuro
	case tcmrsv.CampusNakameguro:
		return valueobject.CampusTypeNakameguro
	default:
		return valueobject.CampusTypeUnknown
	}
}
