package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Robertron624/go-api/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	//Leer el parametro id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Encontrar el libro por el Id

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Eliminar el libro

	h.DB.Delete(book)

	//Enviar un mensaje que se eliminó correcntamente
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
