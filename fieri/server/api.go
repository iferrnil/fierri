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
	tasks := todo.List()
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

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		input := fromReq(r)
		if input == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("input %v", input.ToDo)
		newItem := todo.Add(input.ToDo)
		output := toDoItem(*newItem)
		writeJson(&output, w)
	case http.MethodGet:
		gid := retriveGid(r.URL.Path)
		elem := todo.FindByGid(gid)
		if elem == nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		output := toDoItem(*elem)
		writeJson(&output, w)
	case http.MethodDelete:
	case http.MethodPut:
		input := fromReq(r)
		currentTodo := todo.FindByGid(input.Gid)
		if currentTodo == nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		currentTodo.ToDo = input.ToDo
		updated := todo.Update(*currentTodo)
		output := toDoItem(*updated)
		writeJson(&output, w)
	}
}
