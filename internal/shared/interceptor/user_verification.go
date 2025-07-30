package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/service"
	"github.com/ekkx/tcmrsv-web/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/pkg/actor"
	"github.com/ekkx/tcmrsv-web/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UserVerificationInterceptor(dbPool *pgxpool.Pool) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctxActor := ctxhelper.Actor(ctx)

			if !ctxActor.IsUser() {
				return next(ctx, req)
			}

			userService := service.New(repository.New(database.New(dbPool)))

			user, err := userService.GetUserByID(ctx, ctxActor.ID)
			if err != nil {
				return nil, err
			}

			if user == nil {
				return nil, errs.ErrRequestUserNotFound
			}

			// アクターにそれぞれの役割と公式サイトの認証情報を設定
			if user.IsMaster() {
				ctxActor.WithRole(actor.RoleMaster)
				ctxActor.WithOfficialSiteAuth(&actor.OfficialSiteAuth{
					UserID:   *user.OfficialSiteID,
					Password: *user.OfficialSitePassword, // TODO: 復号化
				})
			} else {
				ctxActor.WithRole(actor.RoleSlave)
				ctxActor.WithOfficialSiteAuth(&actor.OfficialSiteAuth{
					UserID:   *user.MasterUser.OfficialSiteID,
					Password: *user.MasterUser.OfficialSitePassword, // TODO: 復号化
				})
			}

			newCtx := ctxhelper.WithActor(ctx, ctxActor)
			return next(newCtx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
