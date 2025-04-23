package di

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/interface/handler"
)

type Handlers struct {
	User *handler.UserHandler
	// Memo *handler.MemoHandler
}

func NewHandlers(usecases *Usecases) *Handlers {
	return &Handlers{
		User: handler.NewUserHandler(usecases.User),
	}
}

