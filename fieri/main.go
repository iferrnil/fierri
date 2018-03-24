package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/iferrnil/fieri/todo"
)

func serveStaticHandler(fileName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// umieszczenie kodu wyżej powoduje domknięice 'data' -> nie da się odświeżać treści {
		data, err := ioutil.ReadFile("static/" + fileName)
		if err != nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		// }
		w.Write(data)
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
	w.Write(json)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		todo.Add("Test")
	case http.MethodGet:
	case http.MethodDelete:
	case http.MethodPut:
		log.Fatal("Not implemented")
	}
}

func main() {
	wait := make(chan string)
	go func() {
		http.HandleFunc("/", serveStaticHandler("index.html"))
		http.HandleFunc("/testNotFound", serveStaticHandler("not-exists.html"))
		http.HandleFunc("/api/list_task", listTaskHandler)
		http.HandleFunc("/api/task", taskHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
		wait <- "finished"
	}()
	openBrowser("http://localhost:8080")
	<-wait
}

// testowy czy się da - fajna opcja
// source: https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
