// goapi project main.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/rmagnum2002/goapi/models"

	"github.com/gorilla/mux"
)

var books []models.Book

func GetBooksEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func GetBookEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}

func CreateBookEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var book models.Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func DeleteBookEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	fmt.Println("Listening on port 3000")

	books = append(books, models.Book{ID: "1", Title: "Learn Ruby", Content: "some content here for this book", Author: &models.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, models.Book{ID: "2", Title: "Learn Javascript", Content: "some content here for this book", Author: &models.Author{Firstname: "Al", Lastname: "Pacino"}})
	books = append(books, models.Book{ID: "3", Title: "Learn Go", Content: "some content here for this book", Author: &models.Author{Firstname: "Jane", Lastname: "Pane"}})

	//	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	//	http.HandleFunc("/", indexHandler)

	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/books", GetBooksEndpoint).Methods("GET")
	router.HandleFunc("/api/books/{id}", GetBookEndpoint).Methods("GET")
	router.HandleFunc("/api/books", CreateBookEndpoint).Methods("POST")
	router.HandleFunc("/api/books/{id}", GetBooksEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello World")
	t, err := template.ParseFiles("templates/index.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}
