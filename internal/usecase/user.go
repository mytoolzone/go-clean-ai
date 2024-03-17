package usecase

import (
	"context"
	"go-clean/internal/app_err"
	"go-clean/internal/entity"
	"go-clean/pkg/util"
	// Add additional imports if needed for hashing, logging, etc.
)

type userUseCase struct {
	userRepo UserRepo
}

func NewUserUseCase(userRepo UserRepo) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	return uc.userRepo.Create(ctx, user)
}

func (uc *userUseCase) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	return uc.userRepo.FindByID(ctx, id)
}

func (uc *userUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	return uc.userRepo.Update(ctx, user)
}

func (uc *userUseCase) Login(ctx context.Context, mobile, password string) (*entity.User, error) {
	user, err := uc.userRepo.FindByMobile(ctx, mobile)
	if err != nil {
		return nil, app_err.NewError(app_err.ErrLoginFailed, app_err.MsgLoginFailed, err)
	}

	if user.Password != password {
		return nil, app_err.NewError(app_err.ErrLoginFailed, app_err.MsgLoginPasswdErr)
	}
	return user, nil
}

func (uc *userUseCase) Register(ctx context.Context, user *entity.User) error {
	if user.Mobile == "" {
		return app_err.NewError(app_err.ErrBadRequest, app_err.MsgParamNoMobile)
	}

	user.Password = util.Md5(util.Md5(user.Password))
	return uc.CreateUser(ctx, user)
}
