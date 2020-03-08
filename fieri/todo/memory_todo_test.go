package todo

import (
	"testing"
)

func buildTodo() ToDo {
	return NewMemoryTodo(10)
}

func TestAdd(t *testing.T) {
	t.Logf("TestAdd")
	// given
	toDoMan := buildTodo()
	beforeSize := toDoMan.Size()
	// when
	newTodo, err := toDoMan.Add("Whatever")
	if err != nil {
		t.Errorf("Add failed")
	}
	// then
	if toDoMan.Size() != beforeSize+1 {
		t.Errorf("Tasks size should increase, size: %d", toDoMan.Size())
	}
	if elem, _ := toDoMan.FindByGid(newTodo.Gid); elem == nil {
		t.Errorf("Task was not added properly")
	}
}
