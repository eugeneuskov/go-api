package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-api/models"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) Create(listId int, item *models.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	itemId, err := insertTodoItem(tx, item)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err = insertListItems(tx, listId, itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func insertTodoItem(tx *sql.Tx, item *models.TodoItem) (int, error) {
	var id int

	err := tx.QueryRow(
		fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoItemsTable),
		item.Title,
		item.Description,
	).Scan(&id)

	return id, err
}

func insertListItems(tx *sql.Tx, listId, itemId int) error {
	_, err := tx.Exec(
		fmt.Sprintf("insert into %s (list_id, item_id) values ($1, $2)", listItemsTable),
		listId,
		itemId,
	)

	return err
}

func (t *TodoItemPostgres) GetAllByUserList(userId, listId int) ([]models.TodoItem, error) {
	var todoItems []models.TodoItem

	err := t.db.Select(
		&todoItems,
		fmt.Sprintf(
			`select ti.* from %s ti
			inner join %s li on ti.id = li.item_id
			inner join %s ul on li.list_id = ul.list_id
			where li.list_id=$1 and ul.user_id=$2`,
			todoItemsTable,
			listItemsTable,
			usersListsTable,
		),
		listId,
		userId,
	)

	return todoItems, err
}

func (t *TodoItemPostgres) GetById(userId, listId, itemId int) (models.TodoItem, error) {
	var item models.TodoItem

	err := t.db.Get(
		&item,
		fmt.Sprintf(
			`select ti.id, ti.title, ti.description, ti.is_done from %s ti
			inner join %s li on ti.id = li.item_id
			inner join %s ul on li.list_id = ul.list_id
			where ul.user_id=$1 and li.list_id=$2 and ti.id=$3`,
			todoItemsTable,
			listItemsTable,
			usersListsTable,
		),
		userId,
		listId,
		itemId,
	)

	return item, err
}

func (t *TodoItemPostgres) Update(userId, listId, itemId int, input *models.UpdateItemListInput) error {
	setValues := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	argId := 1

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

	if input.IsDone != nil {
		setValues = append(setValues, fmt.Sprintf("is_done=$%b", argId))
		args = append(args, *input.IsDone)
		argId++
	}
	args = append(args, userId, listId, itemId)

	_, err := t.db.Exec(
		fmt.Sprintf(
			`update %s ti set %s from %s li, %s ul
			where ti.id = li.item_id and li.list_id = ul.list_id
			and ul.user_id=$%d and li.list_id=$%d and ti.id=$%d`,
			todoItemsTable,
			strings.Join(setValues, ", "),
			listItemsTable,
			usersListsTable,
			argId,
			argId + 1,
			argId + 2,
		),
		args...
	)

	return err
}

func (t *TodoItemPostgres) DeleteById(userId, listId, itemId int) error {
	_, err := t.db.Exec(
		fmt.Sprintf(
			`delete from %s ti
			using %s li, %s ul
			where ti.id = li.item_id and li.list_id = ul.list_id
			and ul.user_id=$1 and li.list_id=$2 and ti.id=$3`,
			todoItemsTable,
			listItemsTable,
			usersListsTable,
		),
		userId,
		listId,
		itemId,
	)

	return err
}
