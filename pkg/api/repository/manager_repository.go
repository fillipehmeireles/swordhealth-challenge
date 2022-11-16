package repository

import (
	"SwordHealth/pkg/api/drivers"
	dtos "SwordHealth/pkg/api/dtos/manager"
	"SwordHealth/pkg/api/models"
)

type ManagerRepository struct {
	Repository[models.Manager]
}

func NewManagerRepository() *ManagerRepository {
	return &ManagerRepository{Repository[models.Manager]{db: drivers.NewDBDriver()}}
}

func (repo *ManagerRepository) Create(managerDto dtos.CreateManagerDTO) (int, error) {
	repo.db.Connect()
	defer repo.db.Close()
	manager := models.Manager{
		Email:    managerDto.Email,
		Password: managerDto.Password,
		Name:     managerDto.Name,
	}

	result := repo.db.DB.Create(&manager)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil

}

func (repo *ManagerRepository) Update(id int, managerDto dtos.UpdateManagerDTO) (int, error) {
	manager, err := repo.ReadOne(id)
	if err != nil {
		return 0, err
	}

	repo.db.Connect()
	defer repo.db.Close()

	managerModel := models.Manager{
		Email:    managerDto.Email,
		Password: managerDto.Password,
		Name:     managerDto.Name,
	}

	result := repo.db.DB.Model(&manager).Updates(managerModel)

	return int(result.RowsAffected), result.Error
}
