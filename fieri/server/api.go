package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/iferrnil/fieri/todo"
)

type taskAPI struct {
	todo todo.ToDo
}

func NewApi(todo todo.ToDo) *taskAPI {
	return &taskAPI{todo: todo}
}

func handleError(w http.ResponseWriter, err error) {
	log.Fatal("Unexpected api error", err)
	w.WriteHeader(http.StatusInternalServerError)
	return
}

func (t *taskAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	apiTask := strings.TrimPrefix(path, "/api/")
	if strings.HasPrefix(apiTask, "list_task") {
		t.listTaskHandler(w, r)
	} else if strings.HasPrefix(apiTask, "task") {
		t.taskHandler(w, r)
	}
}

func (t *taskAPI) listTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.todo.List()
	if err != nil {
		handleError(w, err)
		return
	}
	output := make([]toDoItem, len(tasks))
	for i, value := range tasks {
		output[i] = toDoItem(value)
	}
	writeJson(output, w)
}

func retriveGid(path string) (gid string) {
	parts := strings.Split(path, "/")
	gid = parts[len(parts)-1]
	return
}

type toDoItem struct {
	ToDo string `json:"todo"`
	Gid  string `json:"gid"`
}

func fromReq(r *http.Request) *toDoItem {
	decoder := json.NewDecoder(r.Body)
	input := &toDoItem{}
	err := decoder.Decode(input)
	if err != nil {
		log.Fatalf("Cannot parse json %v", err)
		return nil
	}
	input.Gid = retriveGid(r.URL.Path)
	return input
}

func writeJson(v interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

func (t *taskAPI) taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		input := fromReq(r)
		if input == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("input %v", input.ToDo)
		newItem, err := t.todo.Add(input.ToDo)
		if err != nil {
			handleError(w, err)
			return
		}
		output := toDoItem(*newItem)
		writeJson(&output, w)
	case http.MethodGet:
		gid := retriveGid(r.URL.Path)
		elem, err := t.todo.FindByGid(gid)
		if err != nil {
			handleError(w, err)
			return
		}
		if elem == nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		output := toDoItem(*elem)
		writeJson(&output, w)
	case http.MethodDelete:
		gid := retriveGid(r.URL.Path)
		deleted, err := t.todo.Remove(gid)
		if err != nil {
			handleError(w, err)
			return
		}
		if deleted == nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		output := toDoItem(*deleted)
		writeJson(output, w)
	case http.MethodPut:
		input := fromReq(r)
		currentTodo, err := t.todo.FindByGid(input.Gid)
		if err != nil {
			handleError(w, err)
			return
		}
		if currentTodo == nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		currentTodo.ToDo = input.ToDo
		updated, err := t.todo.Update(currentTodo)
		if err != nil {
			handleError(w, err)
			return
		}
		output := toDoItem(*updated)
		writeJson(output, w)
	}
}
