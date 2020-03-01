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

type MatchInfo struct {
	Match           bool
	ContentType     string
	FallbackContent string
}

func NewMatch(contentType string, fallback string) MatchInfo {
	return MatchInfo{FallbackContent: fallback, Match: true, ContentType: contentType}
}

func NoMatch() MatchInfo {
	return MatchInfo{Match: false}
}

type ProxyHandler struct {
	PathToFile map[string]string
	ProxyTest  func(string) MatchInfo
	ProxyUrl   string
	ProxyPath  string
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	matchInfo := ph.ProxyTest(path)
	if !matchInfo.Match {
		(&StaticHandler{PathToFile: ph.PathToFile}).ServeHTTP(w, r)
		return
	}
	var client *http.Client = http.DefaultClient
	var fullPath = ph.ProxyUrl + ph.ProxyPath + path
	resp, err := client.Get(fullPath)
	w.Header().Add("Content-Type", matchInfo.ContentType)
	if err != nil {
		// fallback
		w.Write([]byte(matchInfo.FallbackContent))
	} else {
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
