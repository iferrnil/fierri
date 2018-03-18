package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
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

func main() {
	wait := make(chan string)
	go func() {
		http.HandleFunc("/", serveStaticHandler("index.html"))
		http.HandleFunc("/testNotFound", serveStaticHandler("not-exists.html"))
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
