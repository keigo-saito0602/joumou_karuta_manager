package handler

type Handlers struct {
	User *UserHandler
	Memo *MemoHandler
}

func NewHandlers(user *UserHandler, memo *MemoHandler) *Handlers {
	return &Handlers{
		User: user,
		Memo: memo,
	}
}
