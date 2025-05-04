package validation

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/model"
)

type EventScoreValidator struct{}

func NewEventScoreValidator() *EventScoreValidator {
	return &EventScoreValidator{}
}

func (v *EventScoreValidator) CreateEventScoreValidator(ctx context.Context, m *model.EventScoreForCreate) error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Name, validation.Required, validation.RuneLength(1, 20)),
		validation.Field(&m.Feeling, validation.RuneLength(0, 100)),
		validation.Field(&m.CardsTaken, validation.Min(0)),
		validation.Field(&m.FaultCount, validation.Min(0)),
	)
}
