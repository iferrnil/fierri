package todo

import (
	"container/list"
	"log"
)

type MemoryTodo struct {
	GidLen int
	list   *list.List
}

func (mt *MemoryTodo) Add(what string) (res *ToDoItem, err error) {
	res = &ToDoItem{
		ToDo: what,
		Gid:  RandGid(mt.GidLen),
	}
	mt.list.PushFront(res)
	return res, nil
}

func NewMemoryTodo(gidLen int) *MemoryTodo {
	result := &MemoryTodo{GidLen: gidLen, list: list.New()}
	result.init()
	return result
}

func (mt *MemoryTodo) init() {
	mt.Add("Dodać obsługę dodawania przez API")
	mt.Add("Dodać trwałe składowanie")
}

func (mt *MemoryTodo) Update(newValue *ToDoItem) (*ToDoItem, error) {
	for e := mt.list.Front(); e != nil; e = e.Next() {
		elem := e.Value.(*ToDoItem)
		if elem.Gid == newValue.Gid {
			elem.ToDo = newValue.ToDo
			return elem, nil
		}
	}
	return nil, nil
}

func (mt *MemoryTodo) Remove(gid string) (*ToDoItem, error) {
	for e := mt.list.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*ToDoItem)
		if ok && p.Gid == gid {
			copy := ToDoItem(*p)
			mt.list.Remove(e)
			return &copy, nil
		}
	}
	return nil, nil
}

func (mt *MemoryTodo) FindByGid(gid string) (result *ToDoItem, err error) {
	for e := mt.list.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*ToDoItem)
		if ok && p.Gid == gid {
			result = p
			return
		}
	}
	return nil, nil
}

func (mt *MemoryTodo) List() ([]ToDoItem, error) {
	tasks := make([]ToDoItem, mt.list.Len())
	for i, e := 0, mt.list.Front(); e != nil; e, i = e.Next(), i+1 {
		log.Print(e.Value)
		tasks[i] = *(e.Value.(*ToDoItem))
	}
	return tasks, nil
}

func (mt *MemoryTodo) Size() int {
	return mt.list.Len()
}
