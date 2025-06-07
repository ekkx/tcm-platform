package usecase_test

import (
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/usecase"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwter"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestAuthorize_正常系(t *testing.T) {
	mockTCMClient := testhelper.GetMockTCMClient()

	t.Run("新規ユーザー認証", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			output, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         testhelper.TestUserID,
				Password:       testhelper.TestUserPassword,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// トークンの検証
			uID, err := jwter.Verify(output.Authorization.AccessToken, "access", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)
			require.Equal(t, testhelper.TestUserID, uID)

			uID, err = jwter.Verify(output.Authorization.RefreshToken, "refresh", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)
			require.Equal(t, testhelper.TestUserID, uID)

			// ユーザーが存在することを確認
			u, err := userRepo.GetUserByID(ctx, testhelper.TestUserID)
			require.NoError(t, err)
			require.Equal(t, testhelper.TestUserID, u.ID)

			// パスワードが正しく暗号化されていることを確認
			rawPassword, err := cryptohelper.DecryptAES(u.EncryptedPassword, []byte(ctxhelper.GetConfig(ctx).PasswordAESKey))
			require.NoError(t, err)
			require.Equal(t, testhelper.TestUserPassword, rawPassword)
		})
	})

	t.Run("既存ユーザー認証", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// 既存のユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, userRepo)

			output, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         testhelper.TestUserID,
				Password:       testhelper.TestUserPassword,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// トークンの検証
			uID, err := jwter.Verify(output.Authorization.AccessToken, "access", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)
			require.Equal(t, testhelper.TestUserID, uID)

			uID, err = jwter.Verify(output.Authorization.RefreshToken, "refresh", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)
			require.Equal(t, testhelper.TestUserID, uID)
		})
	})
}

func TestAuthorize_異常系(t *testing.T) {
	mockTCMClient := testhelper.GetMockTCMClient()

	t.Run("パラメータの検証エラー", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// ユーザーIDが空の場合
			_, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         "",
				Password:       testhelper.TestUserPassword,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, errs.ErrInvalidArgument)

			// パスワードが空の場合
			_, err = uc.Authorize(ctx, &input.Authorize{
				UserID:         testhelper.TestUserID,
				Password:       "",
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, errs.ErrInvalidArgument)

			// Note:
			// PasswordAESKeyやJWTSecretが空の場合もInvalidArgumentになるが、
			// これはシステム側の設定ミスであるため、将来的にInternalエラーにするべきかもしれない。
		})
	})

	t.Run("ユーザーIDが違う", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			_, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         "wronguser",
				Password:       testhelper.TestUserPassword,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, errs.ErrInvalidEmailOrPassword)

			// ユーザーが新規作成されていないことを確認
			u, err := userRepo.GetUserByID(ctx, "wronguser")
			require.NoError(t, err)
			require.Nil(t, u)
		})
	})

	t.Run("パスワードが違う", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			_, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         testhelper.TestUserID,
				Password:       "wrongpass",
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, errs.ErrInvalidEmailOrPassword)

			// ユーザーが新規作成されていないことを確認
			u, err := userRepo.GetUserByID(ctx, testhelper.TestUserID)
			require.NoError(t, err)
			require.Nil(t, u)
		})
	})
}