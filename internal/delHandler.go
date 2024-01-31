package internal

import "net/http"

func DelHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	delete(Store, key)
	response := []byte("Success")
	w.Write(response)
}
