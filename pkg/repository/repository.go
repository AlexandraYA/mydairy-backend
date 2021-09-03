package repository

import (
	dairy "github.com/alexandraya/mydairy-backend"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user dairy.User) (int, error)
}

type TasksList interface {
}

type Repository struct {
	Authorization
	TasksList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
