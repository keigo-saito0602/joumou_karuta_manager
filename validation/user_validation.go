package validation

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"
)

type UserValidator struct {
	userRepo repository.UserRepository
}

func NewUserValidator(userRepo repository.UserRepository) *UserValidator {
	return &UserValidator{userRepo: userRepo}
}

func (v *UserValidator) ValidateCreate(ctx context.Context, u *model.User) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 100)),
		validation.Field(&u.Email, validation.Required, validation.Length(5, 100), validation.By(ValidateEmailFormat), validation.By(v.EmailAlreadyExistsValidator())),
	)
}

func (v *UserValidator) ValidateUpdate(ctx context.Context, u *model.User) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.ID, validation.Required),
		validation.Field(&u.Name, validation.Required, validation.Length(2, 100)),
		validation.Field(&u.Email, validation.Required, validation.Length(5, 100), validation.By(ValidateEmailFormat)),
	)
}

// UserIDがDBに存在するか
func (v *UserValidator) UserExistsValidator() validation.RuleFunc {
	return func(value interface{}) error {
		userID, ok := value.(uint64)
		if !ok || userID == 0 {
			return nil
		}

		_, err := v.userRepo.GetUser(context.Background(), userID)
		if err != nil {
			return validation.NewError("validation_user", "指定されたユーザーが存在しません")
		}
		return nil
	}
}

// Emailが登録済みかどうか
func (v *UserValidator) EmailAlreadyExistsValidator() validation.RuleFunc {
	return func(value interface{}) error {
		email, ok := value.(string)
		if !ok || email == "" {
			return nil
		}

		_, err := v.userRepo.GetByEmail(context.Background(), email)
		if err == nil {
			return validation.NewError("validation_email", "このメールアドレスは既に使用されています")
		}
		return nil
	}
}
