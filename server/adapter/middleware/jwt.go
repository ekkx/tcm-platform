package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ekkx/tcmrsv-web/server/infra/db"
	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwter"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func JWT(pool *pgxpool.Pool, excludePaths []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// デバッグログ
			path := ctx.Path()
			fmt.Printf("JWT Middleware - Path: %s\n", path)

			// 認証が不必要なパスの除外
			if utils.IsExcludedPath(path, excludePaths) {
				fmt.Printf("JWT Middleware - Path excluded: %s\n", path)
				return next(ctx)
			}

			// 認証トークンの存在をチェック
			auth := ctx.Request().Header.Get("Authorization")
			if auth == "" {
				return apperrors.ErrUnauthorized
			}

			authParts := strings.Split(auth, " ")
			if len(authParts) != 2 || authParts[0] != "Bearer" {
				return apperrors.ErrInvalidToken
			}

			token := authParts[1]
			if token == "" {
				return apperrors.ErrInvalidToken
			}

			cfg := ctxhelper.GetConfig(ctx.Request().Context())

			// 認証トークンを検証
			if uID, err := jwter.Verify(token, "access", []byte(cfg.JWTSecret)); err == nil {
				if uID == nil {
					return apperrors.ErrInvalidToken
				}

				u, err := db.New(pool).GetUserByStudentID(ctx.Request().Context(), *uID)
				if err != nil {
					if errors.Is(err, pgx.ErrNoRows) {
						return apperrors.ErrRequestUserNotFound
					}
					return err
				}

				password, err := cryptohelper.DecryptAES(u.EncryptedPassword, []byte(cfg.PasswordAESKey))
				if err != nil {
					return err
				}

				c := ctxhelper.SetRequestUser(ctx.Request().Context(), ctxhelper.RequestUser{
					ID:       u.ID,
					Password: password,
				})
				ctx.SetRequest(ctx.Request().WithContext(c))
			} else {
				return err
			}

			// 認証トークンをコンテキストに挿入
			c := ctxhelper.SetAccessToken(ctx.Request().Context(), token)
			ctx.SetRequest(ctx.Request().WithContext(c))

			return next(ctx)
		}
	}
}
