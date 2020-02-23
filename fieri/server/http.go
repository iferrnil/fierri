package server

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type StaticHandler struct {
	PathToFile map[string]string
}

type ProxyHandler struct {
	PathToFile   map[string]string
	ProxyPattern func(string) bool
	ProxyUrl     string
	ProxyPath    string
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if !ph.ProxyPattern(path) {
		(&StaticHandler{PathToFile: ph.PathToFile}).ServeHTTP(w, r)
		return
	}
	var client *http.Client = http.DefaultClient
	var fullPath = ph.ProxyUrl + ph.ProxyPath + path
	resp, err := client.Get(fullPath)
	if err != nil {
		// fallback
	}

	log.Print("Server headers:", resp.Header)

	for header, value := range resp.Header {
		if header == "Content-Type" && len(value) > 0 {
			w.Header().Add(header, value[0])
		}
	}

	const BUFFOR_SIZE = 4096
	var buffor []byte = make([]byte, BUFFOR_SIZE)
	for {
		n, err := resp.Body.Read(buffor)
		if n > 0 {
			w.Write(buffor[0:n])
		}
		if io.EOF == err {
			break
		} else if err != nil {
			log.Fatalf("Proxy failed", err)
			break
		}
	}
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
