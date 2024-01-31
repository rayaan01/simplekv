package internal

import (
	"net/http"
)

func SetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	val := r.URL.Query().Get("val")
	Store[key] = val
	response := []byte("Success")
	w.Write(response)
}
