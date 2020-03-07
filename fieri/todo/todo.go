package todo

import (
	"container/list"
	"log"
	"math/rand"
)

var todoList = list.New()

var alfaNumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")

func randGid(length int) string {
	res := make([]rune, length)
	for i := range res {
		res[i] = alfaNumeric[rand.Intn(len(alfaNumeric))]
	}
	return string(res)
}

type ToDoItem struct {
	ToDo string
	Gid  string
}

const gidLen = 10

func Add(what string) (res *ToDoItem) {
	res = &ToDoItem{
		ToDo: what,
		Gid:  randGid(gidLen),
	}
	todoList.PushFront(res)
	return res
}

func init() {
	Add("Dodać obsługę dodawania przez API")
	Add("Dodać trwałe składowanie")
}

func Update(newValue ToDoItem) *ToDoItem {
	for e := todoList.Front(); e != nil; e = e.Next() {
		elem := e.Value.(*ToDoItem)
		if elem.Gid == newValue.Gid {
			elem.ToDo = newValue.ToDo
			return elem
		}
	}
	return nil
}

func Remove(gid string) *ToDoItem {
	for e := todoList.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*ToDoItem)
		if ok && p.Gid == gid {
			copy := ToDoItem(*p)
			todoList.Remove(e)
			return &copy
		}
	}
	return nil
}

func FindByGid(gid string) (result *ToDoItem) {
	for e := todoList.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*ToDoItem)
		if ok && p.Gid == gid {
			result = p
			return
		}
	}
	return nil
}

func List() []ToDoItem {
	tasks := make([]ToDoItem, todoList.Len())
	for i, e := 0, todoList.Front(); e != nil; e, i = e.Next(), i+1 {
		log.Print(e.Value)
		tasks[i] = *(e.Value.(*ToDoItem))
	}
	return tasks
}
