package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personalitie
	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func ShowPersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalitie models.Personalitie
	database.DB.First(&personalitie, id)

	json.NewEncoder(w).Encode(personalitie)
}

func CreatePersonalitie(w http.ResponseWriter, r *http.Request) {
	var personalitie models.Personalitie

	json.NewDecoder(r.Body).Decode(&personalitie)

	database.DB.Create(&personalitie)

	json.NewEncoder(w).Encode(personalitie)
}

func DeletePersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalitie models.Personalitie

	database.DB.Delete(&personalitie, id)

	json.NewEncoder(w).Encode(personalitie)
}

func UpdatePersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalitie models.Personalitie
	database.DB.First(&personalitie, id)

	json.NewDecoder(r.Body).Decode(&personalitie)

	database.DB.Save(&personalitie)

	json.NewEncoder(w).Encode(personalitie)
}
