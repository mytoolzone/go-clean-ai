// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"go-clean/internal/entity/dto"

	"go-clean/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// UserRepo defines the interface for user data storage behavior
	UserRepo interface {
		Create(ctx context.Context, user *entity.User) error
		FindByID(ctx context.Context, id uint) (*entity.User, error)
		FindByOpenID(ctx context.Context, openID string) (*entity.User, error)
		FindByMobile(ctx context.Context, mobile string) (*entity.User, error)
		FindByAPIToken(ctx context.Context, openID string) (*entity.User, error)
		Finds(ctx context.Context, options ...dto.QueryOption) ([]*entity.User, error)
		Update(ctx context.Context, user *entity.User) error
	}
)

type (
	// UserUseCase -.
	UserUseCase interface {
		CreateUser(ctx context.Context, user *entity.User) error
		GetUserByID(ctx context.Context, id uint) (*entity.User, error)
		UpdateUser(ctx context.Context, user *entity.User) error
		Login(ctx context.Context, mobile, password string) (*entity.User, error)
		Register(ctx context.Context, user *entity.User) error
	}

	// Translation -.
	Translation interface {
		History(context.Context) ([]entity.Translation, error)
	}

	// TranslationRepo -.
	TranslationRepo interface {
		GetHistory(context.Context) ([]entity.Translation, error)
	}
)
