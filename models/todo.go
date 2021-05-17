package models

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      string `json:"is_done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
