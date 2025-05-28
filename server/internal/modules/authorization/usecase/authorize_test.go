package usecase_test

import (
	"testing"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/usecase"
	userrepo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/apperrors"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwter"
	mock_tcmrsv "github.com/ekkx/tcmrsv-web/server/tests/mocks/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestAuthorize_正常系(t *testing.T) {
	mockTCMClient := &mock_tcmrsv.MockTCMClient{
		LoginFunc: func(params *tcmrsv.LoginParams) error {
			if params.UserID == "testuser" && params.Password == "testpass" {
				return nil
			}
			return tcmrsv.ErrAuthenticationFailed
		},
	}

	t.Run("新規ユーザー認証", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := userrepo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			output, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         "testuser",
				Password:       "testpass",
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// トークンの検証
			_, err = jwter.Verify(output.Authorization.AccessToken, "access", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)

			_, err = jwter.Verify(output.Authorization.RefreshToken, "refresh", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)

			// ユーザーが存在することを確認
			u, err := userRepo.GetUserByID(ctx, "testuser")
			require.NoError(t, err)
			require.Equal(t, "testuser", u.ID)

			// パスワードが正しく暗号化されていることを確認
			rawPassword, err := cryptohelper.DecryptAES(u.EncryptedPassword, []byte(ctxhelper.GetConfig(ctx).PasswordAESKey))
			require.NoError(t, err)
			require.Equal(t, "testpass", rawPassword)
		})
	})

	t.Run("既存ユーザー認証", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := userrepo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// 既存のユーザーを作成
			_, err := userRepo.CreateUser(ctx, &userrepo.CreateUserArgs{
				ID:                "testuser",
				EncryptedPassword: "encryptedpassword",
			})
			require.NoError(t, err)

			output, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         "testuser",
				Password:       "testpass",
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// トークンの検証
			_, err = jwter.Verify(output.Authorization.AccessToken, "access", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)

			_, err = jwter.Verify(output.Authorization.RefreshToken, "refresh", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)
		})
	})
}

func TestAuthorize_異常系(t *testing.T) {
	mockTCMClient := &mock_tcmrsv.MockTCMClient{
		LoginFunc: func(params *tcmrsv.LoginParams) error {
			if params.UserID == "testuser" && params.Password == "testpass" {
				return nil
			}
			return tcmrsv.ErrAuthenticationFailed
		},
	}

	t.Run("パラメータの検証エラー", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := userrepo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// ユーザーIDが空の場合
			_, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         "",
				Password:       "testpass",
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, apperrors.InvalidArgument)

			// パスワードが空の場合
			_, err = uc.Authorize(ctx, &input.Authorize{
				UserID:         "testuser",
				Password:       "",
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, apperrors.InvalidArgument)

			// Note:
			// PasswordAESKeyやJWTSecretが空の場合もInvalidArgumentになるが、
			// これはシステム側の設定ミスであるため、将来的にInternalエラーにするべきかもしれない。
		})
	})

	t.Run("ユーザーIDが違う", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := userrepo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			_, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         "wronguser",
				Password:       "testpass",
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, apperrors.ErrInvalidEmailOrPassword)

			// ユーザーが新規作成されていないことを確認
			_, err = userRepo.GetUserByID(ctx, "wronguser")
			require.ErrorIs(t, err, apperrors.ErrUserNotFound)
		})
	})

	t.Run("パスワードが違う", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := userrepo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			_, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         "testuser",
				Password:       "wrongpass",
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, apperrors.ErrInvalidEmailOrPassword)

			// ユーザーが新規作成されていないことを確認
			_, err = userRepo.GetUserByID(ctx, "testuser")
			require.ErrorIs(t, err, apperrors.ErrUserNotFound)
		})
	})
}
