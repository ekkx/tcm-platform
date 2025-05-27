package port

import "github.com/ekkx/tcmrsv"

type TCMClient interface {
	Login(params *tcmrsv.LoginParams) error
	GetRoomsFiltered(params tcmrsv.GetRoomsFilteredParams) []tcmrsv.Room
	GetRooms() []tcmrsv.Room
}
