package controllers

import (
	"encoding/json"
	"net/http"
	"go-person-api/models"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {

	pessoa := models.Person{Name: "Fulano", Age: 30}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	err := json.NewEncoder(rw).Encode(pessoa)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}