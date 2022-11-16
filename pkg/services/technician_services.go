package services

import (
	dtos "SwordHealth/pkg/dtos/technician"
	"SwordHealth/pkg/encrypt"
	"SwordHealth/pkg/models"
	"SwordHealth/pkg/repository"
)

type TechnicianServices struct {
	Services[models.Technician, *repository.TechnicianRepository]
}

func NewTechnicianServices() *TechnicianServices {
	return &TechnicianServices{Services[models.Technician, *repository.TechnicianRepository]{repository: repository.NewTechnicianRepository()}}
}

func (service *TechnicianServices) Create(techDto dtos.CreateTechnicianDto) (int, error) {
	hashedPass := encrypt.HashPass(techDto.Password)
	techDto.Password = hashedPass
	return service.repository.Create(techDto)
}

func (service *TechnicianServices) Update(id int, techDto dtos.UpdateTechnicianDto) (int, error) {
	return service.repository.Update(id, techDto)
}
