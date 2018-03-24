package todo

import "container/list"

var TodoList = list.New()

type ToDoItem struct {
	ToDo string `json:"todo"`
}

func Add(what string) {
	TodoList.PushFront(ToDoItem{ToDo: what})
}
