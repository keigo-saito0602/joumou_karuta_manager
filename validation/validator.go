package validation

import "github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"

type Validators struct {
	User       *UserValidator
	Memo       *MemoValidator
	EventScore *EventScoreValidator
}

func NewValidators(userRepo repository.UserRepository) *Validators {
	return &Validators{
		User:       NewUserValidator(userRepo),
		Memo:       NewMemoValidator(NewUserValidator(userRepo)),
		EventScore: NewEventScoreValidator(),
	}
}
