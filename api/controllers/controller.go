package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase"
)

var petRepository = db.NewPetRepository()

func CreateNewPet(w http.ResponseWriter, r *http.Request) {
	var repository = db.NewPetRepository()
	var pet entity.Pet

	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		log.Fatal(err)
	}

	usecase.NewUpdatePetUseCase(repository).Do(pet)

}
