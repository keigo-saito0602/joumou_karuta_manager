package di

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.UserRepository
	Memo repository.MemoRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: repository.NewUserRepository(db),
		Memo: repository.NewMemoRepository(db),
	}
}
