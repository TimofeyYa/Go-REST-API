package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	ListId int `json:"listId"`
}

type TodoItem struct {
	Id          int  `json:"id"`
	Title       int  `json:"title"`
	Description int  `json:"description"`
	Done        bool `json:"done"`
}

type ListsItem struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	ListId int `json:"listId"`
}
