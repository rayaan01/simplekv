package internal

import "net/http"

func GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	response := []byte(Store[key])
	w.Write(response)
}
