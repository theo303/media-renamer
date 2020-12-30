package app

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Starting API...")
	log.Fatal(handleRequests())
}

func handleRequests() error {
	router := mux.NewRouter()
	router.HandleFunc("/", homeLink)
	
	fmt.Println("Starting router on port 3334")
	return http.ListenAndServe(":3334", router)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Media-renamer home")
}