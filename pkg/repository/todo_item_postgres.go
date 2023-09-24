package repository

import (
	"fmt"

	gobackend "github.com/hoach-linux/go-backend"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewItemService(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(item gobackend.CrudTodoItem) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description, done) VALUES ($1, $2, $3) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createListQuery, item.Title, item.Description, "false")

	if err := row.Scan(&id); err != nil {
		tx.Rollback()

		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createUsersListQuery, item.ListId, id)

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
func (r *TodoItemPostgres) GetAll(listId int) ([]gobackend.TodoItem, error) {
	var items []gobackend.TodoItem

	getListQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl INNER JOIN %s ul on tl.id = ul.item_id WHERE ul.list_id = $1", todoItemsTable, listsItemsTable)
	err := r.db.Select(&items, getListQuery, listId)

	return items, err
}
func (r *TodoItemPostgres) GetById(listId, itemId int) (gobackend.TodoItem, error) {
	var item gobackend.TodoItem

	getItemQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl INNER JOIN %s ul on tl.id = ul.item_id WHERE ul.list_id = $1 AND ul.item_id = $2", todoItemsTable, listsItemsTable)
	err := r.db.Get(&item, getItemQuery, listId, itemId)

	return item, err
}
func (r *TodoItemPostgres) Delete(listId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.item_id AND ul.list_id = $1 AND ul.item_id = $2", todoItemsTable, listsItemsTable)
	_, err := r.db.Exec(query, listId, itemId)

	return err
}
