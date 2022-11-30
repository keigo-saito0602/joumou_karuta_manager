package usecase

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"go-clean-api/entity"
	"go-clean-api/util"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var validate *validator.Validate

type loginUseCase struct {
	userRepo entity.UserRepoInterface
	Config   *util.Config
}

func NewLoginUseCase(uRepo entity.UserRepoInterface, config *util.Config) entity.AuthUCaseInterface {
	return &loginUseCase{
		userRepo: uRepo,
		Config:   config,
	}
}

func (l *loginUseCase) ValidateLogin(login *entity.Login) error {
	validate = validator.New()
	err := validate.Struct(login)
	if err != nil {
		return err
	}
	return nil
}

func (l *loginUseCase) CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (l *loginUseCase) GenerateToken(trackUser *entity.User) string {
	claims := &entity.JwtCustomClaims{
		Id:    trackUser.Id,
		Name:  trackUser.Name,
		Email: trackUser.Email,
		Role:  trackUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(l.Config.JwtSecretKey))

	return tokenStr
}

func (l *loginUseCase) Login(lg *entity.Login, bearer *string) error {
	err := l.ValidateLogin(lg)
	if err != nil {
		return err
	}

	trackUser := l.userRepo.GetByUsernameEmail(lg.Username, lg.Password)
	if trackUser == (entity.User{}) {
		return errors.New(fmt.Sprintf("Username or email %s not found", lg.Username))
	}

	isAuth := l.CheckHashPassword(lg.Password, trackUser.Password)
	if isAuth == false {
		return nil
	}

	*bearer = l.GenerateToken(&trackUser)

	return nil
}
