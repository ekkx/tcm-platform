package usecase

import (
	"context"
	"log/slog"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

func (uc *UseCaseImpl) Authorize(ctx context.Context, params *AuthorizeInput) (*AuthorizeOutput, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	userULID, err := ulid.Parse(params.UserID)
	if err == nil && !userULID.IsZero() {
		slog.Debug("slave user login detected", slog.String("user_id", userULID.String()))
		return uc.authorizeByULID(ctx, params, userULID)
	}

	// UserID が ulid ではない場合は公式サイトの UserID と仮定する
	slog.Debug("master user login detected", slog.String("user_id", params.UserID))
	return uc.authorizeByOfficialSite(ctx, params)
}

func (uc *UseCaseImpl) authorizeByOfficialSite(ctx context.Context, params *AuthorizeInput) (*AuthorizeOutput, error) {
	// マスターユーザーの可能性がある場合は、データベースからユーザーを検索
	// 存在していれば、念の為公式サイトにログインしてトークンを生成（公式サイト側のパスワードが変更された可能性があるため）
	// 存在していなければ、公式サイトにログインして新規作成+トークンを生成
	user, err := uc.userService.GetUserByOfficialSiteID(ctx, params.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		slog.Debug("no master user found, attempting to login to official site", slog.String("user_id", params.UserID))

		// 公式サイトにログイン
		if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
			UserID:   params.UserID,
			Password: params.Password,
		}); err != nil {
			slog.Error("failed to login to official site", slog.String("user_id", params.UserID), slog.String("error", err.Error()))
			return nil, errs.ErrUnauthorized
		}

		slog.Debug("official site login successful, creating new master user", slog.String("user_id", params.UserID))

		// ログインできればマスターアカウントを新規作成+トークンを生成
		userID, err := uc.userRepo.CreateUser(ctx, &repository.CreateUserParams{
			Password:             params.Password, // TODO: ハッシュ化する
			OfficialSiteID:       &params.UserID,
			OfficialSitePassword: &params.Password, // TODO: 再利用するため暗号化する
			MasterUserID:         nil,
		})
		if err != nil {
			return nil, err
		}

		slog.Debug("new master user created", slog.String("user_id", userID.String()))

		user, err := uc.userService.GetUserByID(ctx, *userID)
		if err != nil {
			return nil, err
		}

		return uc.issueTokens(user)
	}

	if user.IsMaster() {
		slog.Debug("master user found, attempting to login to official site", slog.String("user_id", *user.OfficialSiteID))
		return uc.handleMasterUserLogin(ctx, params, user)
	}
	return nil, errs.ErrUnauthorized
}

func (uc *UseCaseImpl) authorizeByULID(ctx context.Context, params *AuthorizeInput, id ulid.ULID) (*AuthorizeOutput, error) {
	user, err := uc.userService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil || !user.CheckPassword(params.Password) {
		return nil, errs.ErrUnauthorized
	}

	if user.IsMaster() {
		return uc.handleMasterUserLogin(ctx, params, user)
	}
	return uc.issueTokens(user)
}

func (uc *UseCaseImpl) handleMasterUserLogin(ctx context.Context, params *AuthorizeInput, user *entity.User) (*AuthorizeOutput, error) {
	// 公式サイトにログイン
	if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
		UserID:   params.UserID,
		Password: params.Password,
	}); err != nil {
		return nil, errs.ErrUnauthorized
	}

	slog.Debug("master user login successful", slog.String("user_id", user.ID.String()))

	// TODO: ログインできればパスワードをアップデートしてトークンを生成
	uc.userRepo.UpdateUserByID(ctx, &repository.UpdateUserByIDParams{
		ID:                   user.ID,
		Password:             &params.Password, // TODO: ハッシュ化する
		OfficialSitePassword: &params.Password, // TODO: 再利用するため暗号
	})

	slog.Debug("master user password updated", slog.String("user_id", user.ID.String()), slog.String("official_site_id", *user.OfficialSiteID))

	return uc.issueTokens(user)
}

func (uc *UseCaseImpl) issueTokens(user *entity.User) (*AuthorizeOutput, error) {
	accessToken, err := uc.jwtManager.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := uc.jwtManager.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}
	return NewAuthorizeOutput(entity.Auth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}), nil
}
