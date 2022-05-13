package controllers

import (
	"api/services"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("INFO - GetPokemons - ", r.URL.Path)
	quantity := strings.TrimPrefix(r.URL.Path, "/quantity/")
	data, err := services.RetrievePokemon(quantity)
	if err.Code != 0 {
		log.Println("ERROR - GetPokemons - ", err.Description)
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
