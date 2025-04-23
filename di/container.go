package di

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
	"gorm.io/gorm"
)

type Container struct {
	DB       *gorm.DB
	Repos    *Repositories
	Usecases *Usecases
	Handlers *Handlers
}

func NewContainer() *Container {
	conn := db.NewMySQLDB()
	repos := NewRepositories(conn)
	usecases := NewUsecases(conn, repos)
	handlers := NewHandlers(usecases)

	return &Container{
		DB:       conn,
		Repos:    repos,
		Usecases: usecases,
		Handlers: handlers,
	}
}
