package validation

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
)

type MemoValidator struct {
	userValidator *UserValidator
}

func NewMemoValidator(userValidator *UserValidator) *MemoValidator {
	return &MemoValidator{userValidator: userValidator}
}

func (v *MemoValidator) ValidateCreate(ctx context.Context, m *model.Memo) error {
	return validation.ValidateStructWithContext(ctx, m,
		validation.Field(&m.UserID,
			validation.Required,
			validation.By(validatePositiveUint64),
			validation.By(v.userValidator.UserExistsValidator()),
		),
		validation.Field(&m.Title, validation.Required, validation.RuneLength(2, 20)),
		validation.Field(&m.Active, validation.Required, validation.In(model.BoolFlagFalse, model.BoolFlagTrue)),
		validation.Field(&m.Content, validation.RuneLength(0, 1000)),
	)
}

func (v *MemoValidator) ValidateUpdate(ctx context.Context, m *model.Memo) error {
	return validation.ValidateStructWithContext(ctx, m,
		validation.Field(&m.ID, validation.Required),
		validation.Field(&m.UserID,
			validation.Required,
			validation.By(validatePositiveUint64),
			validation.By(v.userValidator.UserExistsValidator()),
		),
		validation.Field(&m.Title, validation.Required, validation.RuneLength(2, 20)),
		validation.Field(&m.Active, validation.Required, validation.In(model.BoolFlagFalse, model.BoolFlagTrue)),
		validation.Field(&m.Content, validation.RuneLength(0, 1000)),
	)
}
