package main

import (
	"log"
	"net/http"

	"github.com/Robertron624/go-api/pkg/db"
	"github.com/Robertron624/go-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("La API esta corriendo!!")
	http.ListenAndServe(":4000", router)
}
