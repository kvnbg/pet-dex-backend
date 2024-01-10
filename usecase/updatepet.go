package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type updatePetUseCase struct {
	petRepository interfaces.PetRepository
}

func NewUpdatePetUseCase(p interfaces.PetRepository) *updatePetUseCase {
	return &updatePetUseCase{
		petRepository: p,
	}
}

func (auc *updatePetUseCase) Do(petTobeUpdated entity.Pet) {
	fmt.Printf("Pet to be updated %+v\n", petTobeUpdated)
	auc.petRepository.Update(petTobeUpdated)
	fmt.Println("Pet Updated")
}
