package usecase

import (
	"context"

	"github.com/keigo-saito0602/joumou_karuta_manager/auth"
	"github.com/keigo-saito0602/joumou_karuta_manager/config/logger"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	dbctx "github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"
	"github.com/keigo-saito0602/joumou_karuta_manager/util"
	"gorm.io/gorm"
)

type UserUsecase interface {
	GetUser(ctx context.Context, id uint64) (*model.UserResponse, error)
	ListUsers(ctx context.Context) ([]*model.UserResponse, error)
	CreateUser(ctx context.Context, user *model.User) (*model.UserResponse, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.UserResponse, error)
	DeleteUser(ctx context.Context, id uint64) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Login(ctx context.Context, email, password string) (*model.LoginResponse, error)
}

type userUsecase struct {
	db             *gorm.DB
	userRepository repository.UserRepository
}

func NewUserUsecase(db *gorm.DB, repo repository.UserRepository) UserUsecase {
	return &userUsecase{db: db, userRepository: repo}
}

func (u *userUsecase) GetUser(ctx context.Context, id uint64) (*model.UserResponse, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("GetUser called with ID=%d", id)

	user, err := u.userRepository.GetUser(ctx, id)
	if err != nil {
		log.Errorf("failed to get user: %v", err)
		return nil, err
	}

	return ToUserResponse(user), nil
}

func (u *userUsecase) ListUsers(ctx context.Context) ([]*model.UserResponse, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Info("ListUsers called")

	users, err := u.userRepository.ListUsers(ctx)
	if err != nil {
		log.Errorf("failed to list users: %v", err)
		return nil, err
	}

	return ToUserResponseList(users), nil
}

func (u *userUsecase) CreateUser(ctx context.Context, user *model.User) (*model.UserResponse, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("CreateUser called: %+v", user)

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		log.Errorf("failed to hash password: %v", err)
		return nil, domain.WithInternalError("パスワードのハッシュ化に失敗しました")
	}
	user.Password = hashedPassword

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		log.Errorf("failed to create user: %v", err)
		return nil, err
	}

	return ToUserResponse(user), nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *model.User) (*model.UserResponse, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("UpdateUser called: ID=%d", user.ID)

	err := u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		log.Errorf("failed to update user: %v", err)
		return nil, err
	}

	return ToUserResponse(user), nil
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

func (u *userUsecase) Login(ctx context.Context, email, password string) (*model.LoginResponse, error) {
	ctx = dbctx.ToContext(ctx, u.db)
	log := logger.FromContext(ctx)
	log.Infof("Login called with email=%s", email)

	user, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		log.Warnf("email not found: %v", err)
		return nil, domain.WithUnauthenticated("メールアドレスまたはパスワードが間違っています")
	}

	if !util.CheckPasswordHash(password, user.Password) {
		log.Warnf("password mismatch for email=%s", email)
		return nil, domain.WithUnauthenticated("メールアドレスまたはパスワードが間違っています")
	}

	token, err := auth.GenerateJWT(user.ID, user.Role)
	if err != nil {
		log.Errorf("token generation failed: %v", err)
		return nil, domain.WithInternalError("トークン生成に失敗しました")
	}

	return &model.LoginResponse{
		Token: token,
		User:  ToUserResponse(user),
	}, nil
}

func ToUserResponse(u *model.User) *model.UserResponse {
	return &model.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
}

func ToUserResponseList(users []model.User) []*model.UserResponse {
	res := make([]*model.UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, ToUserResponse(&u))
	}
	return res
}
