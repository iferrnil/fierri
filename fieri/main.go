package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/iferrnil/fieri/server"
)

func main() {
	wait := make(chan string)

	taskApi := &server.TaskAPI{}

	var resourceHandler http.Handler = resourceHandler()

	go func() {
		http.Handle("/api/", taskApi)
		http.Handle("/", resourceHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
		wait <- "finished"
	}()
	openBrowser("http://localhost:8080")
	<-wait
}

func resourceHandler() http.Handler {

	devRun := os.Getenv("TODO_DEV_RUN")
	pathToFile := map[string]string{
		"/testNotFound": "not-exists.html",
		"/index.js":     "index.js",
		"/index.js.map": "index.js.map",
		"/":             "index.html",
	}
	if devRun == "" {
		return &server.StaticHandler{
			PathToFile: pathToFile,
		}
	} else {
		proxyPattern := func(path string) bool {
			if strings.HasSuffix(path, ".js") {
				return true
			}
			if strings.HasSuffix(path, ".js.map") {
				return true
			}
			return false
		}
		return &server.ProxyHandler{
			PathToFile:   pathToFile,
			ProxyPattern: proxyPattern,
			ProxyUrl:     "http://localhost:1234",
			ProxyPath:    "",
		}
	}

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
