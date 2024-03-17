package repo

import (
	"context"
	"go-clean/internal/entity"
	"go-clean/internal/entity/dto"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, user *entity.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepo) FindByOpenID(ctx context.Context, openID string) (*entity.User, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).Where("open_id = ?", openID).First(&user).Error
	return &user, err
}

func (r *UserRepo) FindByMobile(ctx context.Context, mobile string) (*entity.User, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).Where("mobile = ?", mobile).First(&user).Error
	return &user, err
}

func (r *UserRepo) FindByAPIToken(ctx context.Context, apiToken string) (*entity.User, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).Where("api_token = ?", apiToken).First(&user).Error
	return &user, err
}

func (r *UserRepo) Finds(ctx context.Context, options ...dto.QueryOption) ([]*entity.User, error) {
	var users []*entity.User
	query := r.DB.WithContext(ctx)
	for _, option := range options {
		option.Apply(query)
	}
	err := query.Find(&users).Error
	return users, err
}

func (r *UserRepo) Update(ctx context.Context, user *entity.User) error {
	return r.DB.WithContext(ctx).Save(user).Error
}
