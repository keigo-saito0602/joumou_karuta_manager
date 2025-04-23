package di

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/interface/handler"
	"github.com/keigo-saito0602/joumou_karuta_manager/validation"
)

type Handlers struct {
	User *handler.UserHandler
	Memo *handler.MemoHandler
	Card *handler.CardHandler
}

func NewHandlers(usecases *Usecases, validators *validation.Validators) *handler.Handlers {
	return &handler.Handlers{
		User: handler.NewUserHandler(usecases.User, validators.User),
		Memo: handler.NewMemoHandler(usecases.Memo, validators.Memo),
		Card: handler.NewCardHandler(usecases.Card),
	}
}
