package main

import (
	"fmt"
	"net/http"
	"singlekv/internal"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/set", internal.SetHandler).Methods("GET")
	router.HandleFunc("/get", internal.GetHandler).Methods("GET")
	router.HandleFunc("/del", internal.DelHandler).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Server listening on 8080")
	http.ListenAndServe(":8080", router)
}
