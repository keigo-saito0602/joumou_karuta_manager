package handler

type Handlers struct {
	User       *UserHandler
	Memo       *MemoHandler
	Card       *CardHandler
	EventScore *EventScoreHandler
}

func NewHandlers(
	user *UserHandler,
	memo *MemoHandler,
	card *CardHandler,
	eventScore *EventScoreHandler,
) *Handlers {
	return &Handlers{
		User:       user,
		Memo:       memo,
		Card:       card,
		EventScore: eventScore,
	}
}
