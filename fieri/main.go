package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/iferrnil/fieri/server"
)

func main() {
	wait := make(chan string)

	taskApi := &server.TaskAPI{}
	pathToFile := map[string]string{
		"/testNotFound": "not-exists.html",
		"/index.js":     "index.js",
		"/index.js.map": "index.js.map",
		"/":             "index.html",
	}
	staticHandler := &server.StaticHandler{
		PathToFile: pathToFile,
	}

	go func() {
		http.Handle("/api/", taskApi)
		http.Handle("/", staticHandler)
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
