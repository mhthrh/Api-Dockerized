package API

import (
	"io"
	"net/http"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func init() {
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		io.WriteString(writer, "<p><h2>Stop killing Ukrainian people for the sake of humanity.</h2></p>")
		return
	}
}

type RequestHandler struct{}

func (*RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	if path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusBadRequest)
		return
	}

	if method, ok := mux[path]; ok {
		method(w, r)
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
	http.Error(w, "I dont know what's happened!!!", http.StatusBadRequest)

}
