package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	gobackend "github.com/hoach-linux/go-backend"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func (r *TodoListPostgres) Create(userId int, list gobackend.TodoList) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()

		return 0, err
	}
}
