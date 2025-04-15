package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/entity"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

var validate *validator.Validate

type userUseCase struct {
	contextTimeout time.Duration
	userRepo       entity.UserRepoInterface
}

func NewUserUseCase(uRepo entity.UserRepoInterface, timeout time.Duration) entity.UserUCaseInterface {

	return &userUseCase{
		userRepo:       uRepo,
		contextTimeout: timeout,
	}
}

func (u *userUseCase) Index(c context.Context) {
	fmt.Println("Print Index")
}

func (u *userUseCase) ValidateInsertUser(ctx context.Context, user *entity.User) error {
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}

	g := new(errgroup.Group)

	g.Go(func() error {
		trackUser := u.userRepo.GetByUsername(user.Username)
		if trackUser != (entity.User{}) {
			return errors.New(fmt.Sprintf("Username %s already taken", user.Username))
		}
		return nil
	})

	g.Go(func() error {
		trackUser := u.userRepo.GetByEmail(user.Email)
		if trackUser != (entity.User{}) {
			return errors.New(fmt.Sprintf("Email %s already taken", user.Email))
		}
		return nil
	})

	g.Go(func() error {
		switch user.Role {
		case "administrator", "user":
			return nil
		default:
			return errors.New(fmt.Sprintf("Role user `%s` is not allowed", user.Role))
		}
	})

	go func() {
		err := g.Wait()
		if err != nil {
			return
		}
	}()

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *userUseCase) InsertUser(ctx context.Context, user *entity.User) error {
	err := u.ValidateInsertUser(ctx, user)
	if err != nil {
		return err
	}

	// hash password
	passwordHash, _ := u.HashPassword(user.Password)
	user.Password = passwordHash

	u.userRepo.StoreUser(user)
	return nil
}
