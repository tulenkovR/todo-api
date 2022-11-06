package todo

type TodoList struct {
	Id          int    `json:"-"`
	Title       string `json:"name"`
	Description string `json:"username"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"name"`
	Description string `json:"username"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
