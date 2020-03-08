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
	beforeSize, _ := toDoMan.Size()
	// when
	newTodo, err := toDoMan.Add("Whatever")
	if err != nil {
		t.Errorf("Add failed")
	}
	// then
	if afterSize, _ := toDoMan.Size(); afterSize != beforeSize+1 {
		t.Errorf("Tasks size should increase, size: %d", afterSize)
	}
	if elem, _ := toDoMan.FindByGid(newTodo.Gid); elem == nil {
		t.Errorf("Task was not added properly")
	}
}
