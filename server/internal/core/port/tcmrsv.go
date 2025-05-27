package port

import "github.com/ekkx/tcmrsv"

type TCMClient interface {
	Login(params *tcmrsv.LoginParams) error
}
