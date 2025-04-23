package repository

import (
	"context"
	"errors"
	"log"

	"github.com/keigo-saito0602/joumou_karuta_manager/domain"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	ListUsers(ctx context.Context) ([]model.User, error)
	GetUser(ctx context.Context, id uint64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uint64) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) CreateUser(ctx context.Context, user *model.User) error {

	err := r.db.Create(user).Error
	if err != nil {
		log.Printf("[UserRepository][CreateUser] Failed to create user: %v", err)
	}
	return err
}

func (r *gormUserRepository) ListUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		log.Printf("[UserRepository][ListUsers] Failed to fetch users: %v", err)
	}
	return users, err
}

func (r *gormUserRepository) GetUser(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		log.Printf("[UserRepository][GetUser] Failed to find user with ID=%d: %v", id, err)
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.User{}).
			Where("id = ?", user.ID).
			Updates(map[string]interface{}{
				"name":  user.Name,
				"email": user.Email,
			}).Error; err != nil {
			log.Printf("[UserRepository][UpdateUser] Failed to update user ID=%d: %v", user.ID, err)
			return err
		}

		if err := tx.First(user, user.ID).Error; err != nil {
			log.Printf("[UserRepository][UpdateUser] Failed to refetch updated user ID=%d: %v", user.ID, err)
			return err
		}

		return nil
	})
}

func (r *gormUserRepository) DeleteUser(ctx context.Context, id uint64) error {
	err := r.db.Delete(&model.User{}, id).Error
	if err != nil {
		log.Printf("[UserRepository][DeleteUser] Failed to delete user ID=%d: %v", id, err)
	}
	return err
}

func (r *gormUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
