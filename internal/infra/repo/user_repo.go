package repo

import (
	"context"
	"golang-ddd/internal/domain"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) Update(ctx context.Context, user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) Delete(ctx context.Context, id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *UserRepo) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id)
	return &user, result.Error
}
