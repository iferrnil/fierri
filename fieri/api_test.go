package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/iferrnil/fieri/server"
	"github.com/iferrnil/fieri/todo"
)

func buildApi() http.Handler {
	var todo todo.ToDo = todo.NewMemoryTodo(10)
	return server.NewApi(todo)
}

func TestGetTasks(t *testing.T) {
	t.Logf("TestGetTasks")
	taskApi := buildApi()
	request, err := http.NewRequest("GET", "/api/list_task", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.Handler(taskApi)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var v []map[string]string
	response := recorder.Body.String()
	t.Logf("Reponse %v", response)
	decoder := json.NewDecoder(strings.NewReader(response))
	jsonParseErr := decoder.Decode(&v)
	if jsonParseErr != nil {
		t.Errorf("Invalid json returned %v", jsonParseErr)
	}
	if len(v) != 2 {
		t.Errorf("By default list results 2 elemnts")
	}
}

func TestAddAndGetTask(t *testing.T) {
	t.Logf("TestAddAndGetTask")
	var taskApi http.Handler = buildApi()
	createRequest, cErr := http.NewRequest("POST", "/api/task", strings.NewReader("{\"todo\": \"Test\"}"))
	if cErr != nil {
		t.Fatal(cErr)
	}
	createRecorder := httptest.NewRecorder()
	handler := http.Handler(taskApi)
	handler.ServeHTTP(createRecorder, createRequest)

	if status := createRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var v map[string]string
	response := createRecorder.Body.String()
	t.Logf("Reponse %v", response)
	decoder := json.NewDecoder(strings.NewReader(response))
	jsonParseErr := decoder.Decode(&v)
	if jsonParseErr != nil {
		t.Errorf("Invalid json returned %v", jsonParseErr)
	}
	if v["todo"] != "Test" {
		t.Errorf("Api should return what was created")
	}
	gid := v["gid"]
	if gid == "" {
		t.Errorf("Result should contains a gid")
	}

	getRequest, gErr := http.NewRequest("GET", "/api/task/"+gid, nil)
	if gErr != nil {
		t.Fatal(gErr)
	}
	getRecorder := httptest.NewRecorder()
	handler.ServeHTTP(getRecorder, getRequest)

}
