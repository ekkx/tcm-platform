package usecase_test

import (
	"testing"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/usecase"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwter"
	mock_tcmrsv "github.com/ekkx/tcmrsv-web/server/tests/mocks/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestReauthorize_正常系(t *testing.T) {
	testUserID := "testuser"
	testPassword := "testpass"
	mockTCMClient := &mock_tcmrsv.MockTCMClient{
		LoginFunc: func(params *tcmrsv.LoginParams) error {
			if params.UserID == testUserID && params.Password == testPassword {
				return nil
			}
			return tcmrsv.ErrAuthenticationFailed
		},
	}

	t.Run("再認証", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// 初回認証でユーザーを作成
			output, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         testUserID,
				Password:       testPassword,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// アクセストークンとリフレッシュトークンが生成されていることを確認
			output2, err := uc.Reauthorize(ctx, &input.Reauthorize{
				RefreshToken:   output.Authorization.RefreshToken,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// トークンの検証
			uID, err := jwter.Verify(output2.Authorization.AccessToken, "access", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)
			require.Equal(t, testUserID, uID)

			uID, err = jwter.Verify(output2.Authorization.RefreshToken, "refresh", []byte(ctxhelper.GetConfig(ctx).JWTSecret))
			require.NoError(t, err)
			require.Equal(t, testUserID, uID)
		})
	})
}

func TestReauthorize_異常系(t *testing.T) {
	testUserID := "testuser"
	testPassword := "testpass"
	mockTCMClient := &mock_tcmrsv.MockTCMClient{
		LoginFunc: func(params *tcmrsv.LoginParams) error {
			if params.UserID == testUserID && params.Password == testPassword {
				return nil
			}
			return tcmrsv.ErrAuthenticationFailed
		},
	}

	t.Run("リフレッシュトークンのスコープが不正", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// 初回認証でユーザーを作成
			output, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         testUserID,
				Password:       testPassword,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// アクセストークンで再認証を試みる
			_, err = uc.Reauthorize(ctx, &input.Reauthorize{
				RefreshToken:   output.Authorization.AccessToken,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, errs.ErrInvalidJWTScope)
		})
	})

	t.Run("リフレッシュトークンが期限切れ", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// 初回認証でユーザーを作成
			_, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         testUserID,
				Password:       testPassword,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// 期限切れのリフレッシュトークンを生成
			refreshToken, err := jwter.Generate(
				jwt.MapClaims{
					"sub":   testUserID,
					"exp":   jwt.NewNumericDate(time.Now().Add(-24 * time.Hour)),
					"scope": "refresh",
				},
				[]byte(ctxhelper.GetConfig(ctx).JWTSecret),
			)
			require.NoError(t, err)

			// 期限切れのリフレッシュトークンで再認証を試みる
			_, err = uc.Reauthorize(ctx, &input.Reauthorize{
				RefreshToken:   refreshToken,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, errs.ErrRefreshTokenExpired)
		})
	})

	t.Run("ユーザーが存在しない", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// 存在しないユーザーのリフレッシュトークンを生成
			refreshToken, err := jwter.Generate(
				jwt.MapClaims{
					"sub":   "nonexistentuser",
					"exp":   jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
					"scope": "refresh",
				},
				[]byte(ctxhelper.GetConfig(ctx).JWTSecret),
			)
			require.NoError(t, err)

			// 存在しないユーザーで再認証を試みる
			_, err = uc.Reauthorize(ctx, &input.Reauthorize{
				RefreshToken:   refreshToken,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, errs.ErrRequestUserNotFound)
		})
	})

	t.Run("TCMのパスワードが変更された", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			userRepo := user_repo.NewRepository(db)
			uc := usecase.NewUsecase(mockTCMClient, userRepo)

			// 初回認証でユーザーを作成
			output, err := uc.Authorize(ctx, &input.Authorize{
				UserID:         testUserID,
				Password:       testPassword,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.NoError(t, err)

			// ユーザーのパスワードを不正なものに更新
			encryptedPassword, err := cryptohelper.EncryptAES("wrongpass", []byte(ctxhelper.GetConfig(ctx).PasswordAESKey))
			require.NoError(t, err)

			userRepo.UpdateUserPassword(ctx, &user_repo.UpdateUserPasswordArgs{
				ID:                testUserID,
				EncryptedPassword: encryptedPassword,
			})

			// 不正なパスワードで再認証を試みる
			_, err = uc.Reauthorize(ctx, &input.Reauthorize{
				RefreshToken:   output.Authorization.RefreshToken,
				PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
				JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
			})
			require.ErrorIs(t, err, errs.ErrInvalidEmailOrPassword)
		})
	})
}
