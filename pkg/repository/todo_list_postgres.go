package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	gobackend "github.com/hoach-linux/go-backend"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListProgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
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

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)

	if err != nil {
		tx.Rollback()

		return 0, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}
func (r *TodoListPostgres) GetAll(userId int) ([]gobackend.TodoList, error) {
	var lists []gobackend.TodoList

	getListQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, usersListsTable)
	err := r.db.Select(&lists, getListQuery, userId)

	return lists, err
}
func (r *TodoListPostgres) GetById(userId, listId int) (gobackend.TodoList, error) {
	var list gobackend.TodoList

	getListQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)
	err := r.db.Get(&list, getListQuery, userId, listId)

	return list, err
}
func (r *TodoListPostgres) Delete(userId, listId int) (error) {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

