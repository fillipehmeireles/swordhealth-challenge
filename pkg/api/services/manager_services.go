package services

import (
	dtos "SwordHealth/pkg/api/dtos/manager"
	"SwordHealth/pkg/api/encrypt"
	"SwordHealth/pkg/api/models"
	"SwordHealth/pkg/api/repository"
)

type ManagerServices struct {
	Services[models.Manager, *repository.ManagerRepository]
}

func NewManagerServices() *ManagerServices {
	return &ManagerServices{Services[models.Manager, *repository.ManagerRepository]{repository: repository.NewManagerRepository()}}
}

func (service *ManagerServices) Create(managerDto dtos.CreateManagerDTO) (int, error) {
	hashedPass := encrypt.HashPass(managerDto.Password)
	managerDto.Password = hashedPass
	return service.repository.Create(managerDto)
}

func (service *ManagerServices) Update(id int, managerDto dtos.UpdateManagerDTO) (int, error) {
	return service.repository.Update(id, managerDto)
}
