package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func AllPersonas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func OnePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, persona := range models.Personalidades {
		if strconv.Itoa(persona.Id) == id {
			json.NewEncoder(w).Encode(persona)
		}
	}
}
