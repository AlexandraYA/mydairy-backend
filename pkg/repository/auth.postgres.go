package repository

import (
	dairy "github.com/alexandraya/mydairy-backend"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user dairy.User) (int, error) {

	return 1, nil
}
