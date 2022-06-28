package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Robertron624/go-api/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	//Leer la id del libro
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Leer el cuerpo de la peticion
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateBook models.Book
	json.Unmarshal(body, &updateBook)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = updateBook.Title
	book.Author = updateBook.Author
	book.Desc = updateBook.Desc

	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
