package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gabriel-Rossi-dev/API--POSTGRESQL/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
