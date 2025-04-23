package usecase

import (
	"context"

	"github.com/keigo-saito0602/joumou_karuta_manager/config/logger"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	dbctx "github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"
	"github.com/keigo-saito0602/joumou_karuta_manager/util"
	"gorm.io/gorm"
)

type UserUsecase interface {
	GetUser(ctx context.Context, id uint64) (*model.User, error)
	ListUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uint64) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type userUsecase struct {
	db             *gorm.DB
	userRepository repository.UserRepository
}

func NewUserUsecase(db *gorm.DB, repo repository.UserRepository) UserUsecase {
	return &userUsecase{db: db, userRepository: repo}
}

func (u *userUsecase) GetUser(ctx context.Context, id uint64) (*model.User, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("GetUser called with ID=%d", id)

	user, err := u.userRepository.GetUser(ctx, id)
	if err != nil {
		log.Errorf("failed to get user: %v", err)
	}
	return user, err
}

func (u *userUsecase) ListUsers(ctx context.Context) ([]model.User, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Info("ListUsers called")

	users, err := u.userRepository.ListUsers(ctx)
	if err != nil {
		log.Errorf("failed to list users: %v", err)
	}
	return users, err
}

func (u *userUsecase) CreateUser(ctx context.Context, user *model.User) error {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("CreateUser called: %+v", user)

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		log.Errorf("failed to hash password: %v", err)
		return domain.WithInternalError("パスワードのハッシュ化に失敗しました")
	}
	user.Password = hashedPassword

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		log.Errorf("failed to create user: %v", err)
		return err
	}

	return nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *model.User) error {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("UpdateUser called: ID=%d", user.ID)

	err := u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		log.Errorf("failed to update user: %v", err)
	}
	return err
}

func (u *userUsecase) DeleteUser(ctx context.Context, id uint64) error {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("DeleteUser called: ID=%d", id)

	err := u.userRepository.DeleteUser(ctx, id)
	if err != nil {
		log.Errorf("failed to delete user: %v", err)
	}
	return err
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("GetByEmail called with ID=%s", email)

	user, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		log.Errorf("failed to get user: %v", err)
	}
	return user, err
}
