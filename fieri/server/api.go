package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/iferrnil/fieri/todo"
)

type TaskAPI struct {
}

func (t *TaskAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	apiTask := strings.TrimPrefix(path, "/api/")
	if strings.HasPrefix(apiTask, "list_task") {

		listTaskHandler(w, r)
	} else if strings.HasPrefix(apiTask, "task") {
		taskHandler(w, r)
	}
}

func listTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks = make([]todo.ToDoItem, todo.TodoList.Len())
	for i, e := 0, todo.TodoList.Front(); e != nil; e, i = e.Next(), i+1 {
		log.Print(e.Value)
		tasks[i] = e.Value.(todo.ToDoItem)
	}
	log.Print(tasks)
	json, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

func retriveGid(path string) (gid string) {
	parts := strings.Split(path, "/")
	gid = parts[len(parts)-1]
	return
}

type taskInput struct {
	Data string `json:"todo"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		input := &taskInput{}
		decoder.Decode(input)
		todo.Add(input.Data)
		// zwróc listę
		listTaskHandler(w, r)
	case http.MethodGet:
		gid := retriveGid(r.URL.Path)
		todo := todo.FindByGid(gid)
		if todo == nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		json, err := json.Marshal(todo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	case http.MethodDelete:
	case http.MethodPut:
		log.Fatal("Not implemented")
	}
}
