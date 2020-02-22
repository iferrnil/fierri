package server

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type StaticHandler struct {
	PathToFile map[string]string
}

func (sh *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fileName, ok := sh.PathToFile[path]
	if !ok {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	// umieszczenie kodu wyżej powoduje domknięice 'data' -> nie da się odświeżać treści {
	data, err := ioutil.ReadFile("static/" + fileName)
	if err != nil {
		http.Error(w, "No expected resource, fileName: "+fileName, 500)
		return
	}
	contentType := "text/plain"
	if strings.HasSuffix(fileName, ".js") {
		contentType = "text/javascript"
	} else if strings.HasSuffix(fileName, ".html") {
		contentType = "text/html"
	}
	w.Header().Add("Content-Type", contentType+"; charset=UTF-8")
	w.Write(data)
}
