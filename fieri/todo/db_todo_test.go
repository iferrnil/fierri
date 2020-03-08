package todo

import (
	"context"
	"testing"

	"github.com/iferrnil/fieri/database"
)

func connectToDb(t *testing.T) *DbTodo {
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("Cannot connect to DB %v", err)
	}
	ctx := context.Background()
	return NewDbTodo(10, ctx, db)
}

func TestAddDB(t *testing.T) {
	// given
	toDoMan := connectToDb(t)
	defer toDoMan.db.Close()
	beforeSize, err := toDoMan.Size()
	if err != nil {
		t.Fatalf("Size fetch failed %v", err)
	}
	// when
	newTodo, err := toDoMan.Add("Whatever")
	if err != nil {
		t.Fatalf("Add failed %v", err)
	}
	// then
	afterSize, err := toDoMan.Size()
	if err != nil {
		t.Fatalf("Size fetch failed %v", err)
	}
	if afterSize != beforeSize+1 {
		t.Fatalf("Tasks size should increase, size: %d", afterSize)
	}
	elem, err := toDoMan.FindByGid(newTodo.Gid)
	if err != nil {
		t.Fatalf("Cannot find added element %v", err)
	}
	if elem == nil {
		t.Fatalf("Task was not added properly")
	}
	// czyscimy
	if _, err := toDoMan.Remove(elem.Gid); err != nil {
		t.Fatalf("Cannot clean after test, please remove by hand")
	}
}
