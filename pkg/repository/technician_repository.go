package repository

import (
	"SwordHealth/pkg/drivers"
	dtos "SwordHealth/pkg/dtos/technician"
	"SwordHealth/pkg/models"
)

type TechnicianRepository struct {
	Repository[models.Technician]
}

func NewTechnicianRepository() *TechnicianRepository {
	return &TechnicianRepository{Repository[models.Technician]{db: drivers.NewDBDriver()}}
}

func (repo *TechnicianRepository) Create(techDto dtos.CreateTechnicianDto) (int, error) {
	repo.db.Connect()
	defer repo.db.Close()

	tech := models.Technician{
		Email:    techDto.Email,
		Password: techDto.Password,
		Name:     techDto.Name,
	}

	result := repo.db.DB.Create(&tech)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *TechnicianRepository) Update(id int, techDto dtos.UpdateTechnicianDto) (int, error) {
	tech, err := repo.ReadOne(id)
	if err != nil {
		return 0, err
	}

	repo.db.Connect()
	defer repo.db.Close()

	techModel := models.Technician{
		Email:    techDto.Email,
		Password: techDto.Password,
		Name:     techDto.Name,
	}

	result := repo.db.DB.Model(&tech).Updates(techModel)

	return int(result.RowsAffected), result.Error
}

func (repo *TechnicianRepository) Delete(id int) (int, error) {
	tech, err := repo.ReadOne(id)
	if err != nil {
		return 0, err
	}

	repo.db.Connect()
	defer repo.db.Close()

	result := repo.db.DB.Delete(tech)

	return int(result.RowsAffected), result.Error
}
