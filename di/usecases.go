package di

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/usecase"
	"gorm.io/gorm"
)

type Usecases struct {
	User usecase.UserUsecase
	Memo usecase.MemoUsecase
	Card usecase.CardUsecase
}

func NewUsecases(db *gorm.DB, repository *Repositories) *Usecases {
	return &Usecases{
		User: usecase.NewUserUsecase(db, repository.User),
		Memo: usecase.NewMemoUsecase(db, repository.Memo),
		Card: usecase.NewCardUsecase(db, repository.Card),
	}
}
