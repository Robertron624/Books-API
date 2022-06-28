package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Robertron624/go-api/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	//Leer el cuerpo de la Request
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	//Agregar a tabla de Books
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Enviar una respuesta 201 para indicar que se agrego correctamente
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
