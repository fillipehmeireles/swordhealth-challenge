package services

import (
	dtos "SwordHealth/pkg/api/dtos/technician"
	"SwordHealth/pkg/api/encrypt"
	"SwordHealth/pkg/api/models"
	"SwordHealth/pkg/api/repository"
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
