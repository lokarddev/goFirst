package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"goFirst"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user goFirst.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, username, password_hash) values ($1, $2, $3) returning id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	fmt.Println(query)
	if err := row.Scan(&id); err != nil {
		return 9, nil
	}
	return id, nil
}
