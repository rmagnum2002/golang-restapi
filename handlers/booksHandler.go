package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rmagnum2002/goapi/models"

	"github.com/gorilla/mux"
)

var books []models.Book

func init() {
	books = append(books, models.Book{ID: "1", Title: "Learn Ruby", Content: "some content here for this book", Author: &models.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, models.Book{ID: "2", Title: "Learn Javascript", Content: "some content here for this book", Author: &models.Author{Firstname: "Al", Lastname: "Pacino"}})
	books = append(books, models.Book{ID: "3", Title: "Learn Go", Content: "some content here for this book", Author: &models.Author{Firstname: "Jane", Lastname: "Pane"}})
}

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
