package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Robertron624/go-api/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	//Leer el parametro Id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Encontrar libro por Id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Si los Ids son iguales envia como respuesta la info del libro
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)

}
