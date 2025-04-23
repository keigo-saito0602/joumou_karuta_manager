package di

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/usecase"
	"gorm.io/gorm"
)

type Usecases struct {
	User usecase.UserUsecase
	// Memo usecase.MemoUsecase
}

func NewUsecases(db *gorm.DB, repo *Repositories) *Usecases {
	return &Usecases{
		User: usecase.NewUserUsecase(db, repo.User),
		// Memo: usecase.NewMemoUsecase(db, repo.Memo),
	}
}
