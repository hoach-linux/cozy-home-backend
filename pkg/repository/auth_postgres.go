package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	gobackend "github.com/hoach-linux/go-backend"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func newAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user gobackend.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}
func (r *AuthPostgres) GetUser(username, password string) (gobackend.User, error) {
	var user gobackend.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
