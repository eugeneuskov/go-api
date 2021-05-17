package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-api/models"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list *models.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	todoId, err := insertTodoList(tx, list)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err = insertUserList(tx, userId, todoId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return todoId, tx.Commit()
}

func insertTodoList(tx *sql.Tx, list *models.TodoList) (int, error) {
	var id int

	err := tx.QueryRow(
		fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoListsTable),
		list.Title,
		list.Description,
	).Scan(&id)

	return id, err
}

func insertUserList(tx *sql.Tx, userId int, todoId int) error {
	_, err := tx.Exec(
		fmt.Sprintf("insert into %s (user_id, list_id) values ($1, $2)", usersListsTable),
		userId,
		todoId,
	)

	return err
}