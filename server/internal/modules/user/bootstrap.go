package user

import "github.com/ekkx/tcmrsv-web/server/internal/modules/user/handler"

func InitModule() *handler.Handler {
    return handler.New()
}
