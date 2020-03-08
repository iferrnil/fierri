package todo

import (
	"math/rand"
)

var alfaNumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")

func RandGid(length int) string {
	res := make([]rune, length)
	for i := range res {
		res[i] = alfaNumeric[rand.Intn(len(alfaNumeric))]
	}
	return string(res)
}

type ToDo interface {
	FindByGid(string) (*ToDoItem, error)
	Add(string) (*ToDoItem, error)
	Remove(string) (*ToDoItem, error)
	Update(*ToDoItem) (*ToDoItem, error)
	List() ([]ToDoItem, error)
}

type ToDoItem struct {
	ToDo string
	Gid  string
}
