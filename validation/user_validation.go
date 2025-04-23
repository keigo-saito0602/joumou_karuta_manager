package validation

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
)

type UserValidator struct{}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (v *UserValidator) ValidateCreate(ctx context.Context, u *model.User) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 100)),
		validation.Field(&u.Email, validation.Required, validation.Length(5, 100), validation.By(ValidateEmailFormat)),
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
func (v *MemoValidator) userExistsValidator() validation.RuleFunc {
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
