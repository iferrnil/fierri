package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	wait := make(chan string)
	go func() {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8080", nil))
		wait <- "finished"
	}()
	openBrowser("http://localhost:8080")
	<-wait
}

//source: https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
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
