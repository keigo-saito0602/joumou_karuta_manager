package validation

type Validators struct {
	User *UserValidator
}

func NewValidators() *Validators {
	return &Validators{
		User: NewUserValidator(),
	}
}
