package todo

import "container/list"
import "math/rand"

var TodoList = list.New()

var alfaNumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")

func randGid(length int) string {
	res := make([]rune, length)
	for i := range res {
		res[i] = alfaNumeric[rand.Intn(len(alfaNumeric))]
	}
	return string(res)
}

type ToDoItem struct {
	ToDo string `json:"todo"`
	Gid  string `json:"gid"`
}

const gidLen = 10

func Add(what string) {
	TodoList.PushFront(ToDoItem{
		ToDo: what,
		Gid:  randGid(gidLen),
	})
}

func FindByGid(gid string) (result *ToDoItem) {
	for e := TodoList.Front(); e != nil; e = e.Next() {
		if e.Value.(ToDoItem).Gid == gid {
			p, ok := e.Value.(ToDoItem)
			if ok {
				result = &p
				return
			}
		}
	}
	return nil
}
