package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/iferrnil/iferrnil/fieri/todo"
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
		contentType := "text/plain"
		if strings.HasSuffix(fileName, ".js") {
			contentType = "text/javascript"
		} else if strings.HasSuffix(fileName, ".html") {
			contentType = "text/html"
		}
		w.Header().Add("Content-Type", contentType+"; charset=UTF-8")
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

func main() {
	wait := make(chan string)
	go func() {
		http.HandleFunc("/testNotFound", serveStaticHandler("not-exists.html"))
		http.HandleFunc("/api/list_task", listTaskHandler)
		http.HandleFunc("/api/task/", taskHandler)
		http.HandleFunc("/api/task", taskHandler)
		http.HandleFunc("/index.js", serveStaticHandler("index.js"))
		http.HandleFunc("/index.js.map", serveStaticHandler("index.js.map"))
		http.HandleFunc("/", serveStaticHandler("index.html"))
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
