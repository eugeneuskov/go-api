package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-api/models"
	"strings"
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

func (r *TodoListPostgres) GetAllByUser(userId int) ([]models.TodoList, error) {
	var todoLists []models.TodoList

	err := r.db.Select(
		&todoLists,
		fmt.Sprintf(
			"select tl.* from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id=$1",
			todoListsTable,
			usersListsTable,
		),
		userId,
	)

	return todoLists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (models.TodoList, error) {
	var list models.TodoList

	err := r.db.Get(
		&list,
		fmt.Sprintf(
			`select tl.id, tl.title, tl.description
				from %s tl 
				inner join %s ul on tl.id = ul.list_id 
				where ul.user_id=$1 and ul.list_id=$2`,
			todoListsTable,
			usersListsTable,
		),
		userId,
		listId,
	)

	return list, err
}

func (r *TodoListPostgres) DeleteById(userId, listId int) error {
	_, err := r.db.Exec(
		fmt.Sprintf(
			`delete from %s tl
				using %s ul
				where tl.id = ul.list_id and ul.user_id=$1 and ul.list_id=$2`,
			todoListsTable,
			usersListsTable,
		),
		userId,
		listId,
	)

	return err
}

func (r *TodoListPostgres) Update(userId, listId int, input *models.UpdateListInput) error {
	setValues := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)
	argId := 1;

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	args = append(args, listId, userId)

	_, err := r.db.Exec(
		fmt.Sprintf(
			`update %s tl
			set %s from %s ul
			where tl.id = ul.list_id and ul.list_id=$%d and ul.user_id=$%d`,
			todoListsTable,
			strings.Join(setValues, ", "),
			usersListsTable,
			argId,
			argId + 1,
		),
		args...
	)

	return err
}
