package handler

type Handlers struct {
	User *UserHandler
	Memo *MemoHandler
	Card *CardHandler
}

func NewHandlers(
	user *UserHandler,
	memo *MemoHandler,
	card *CardHandler,
) *Handlers {
	return &Handlers{
		User: user,
		Memo: memo,
		Card: card,
	}
}
