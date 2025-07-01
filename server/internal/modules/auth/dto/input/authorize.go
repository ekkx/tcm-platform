package input

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	authv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type Authorize struct {
	UserID   ulid.ULID
	Password string
}

func NewAuthorize() *Authorize {
	return &Authorize{}
}

func (st *Authorize) Validate() error {
	if st.UserID.IsZero() {
		return errs.ErrInvalidIDFormat
	}
	return nil
}

func (st *Authorize) FromRequest(req *connect.Request[authv1.AuthorizeRequest]) *Authorize {
	parsedID, err := ulid.Parse(req.Msg.UserId)
	if err != nil {
		parsedID = ulid.ULID{}
	}
	return &Authorize{
		UserID:   parsedID,
		Password: req.Msg.Password,
	}
}
