package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	urlshortener "github.com/ramil66/url-shortener"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user urlshortener.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email,passwordHash) values ($1,$2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (urlshortener.User, error) {
	var user urlshortener.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND passwordHash=$2", userTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
