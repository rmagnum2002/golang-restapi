// goapi project main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rmagnum2002/goapi/handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Listening on port 3000")
	//	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	//	http.HandleFunc("/", indexHandler)

	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/api/books", handlers.GetBooksEndpoint).Methods("GET")
	router.HandleFunc("/api/books/{id}", handlers.GetBookEndpoint).Methods("GET")
	router.HandleFunc("/api/books", handlers.CreateBookEndpoint).Methods("POST")
	router.HandleFunc("/api/books/{id}", handlers.GetBooksEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
