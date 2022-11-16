package repository

import (
	"SwordHealth/pkg/drivers"
)

type IRepository[T any] interface {
	ReadAll() ([]T, error)
	ReadOne(id int) (T, error)
	Delete(id int) (int, error)
}

type Repository[T interface{}] struct {
	db *drivers.DatabaseDriver
}

func (repo *Repository[T]) ReadAll() ([]T, error) {
	repo.db.Connect()
	defer repo.db.Close()

	var allData []T
	repo.db.DB.Find(&allData)
	return allData, nil
}

func (repo *Repository[T]) ReadOne(id int) (T, error) {
	repo.db.Connect()
	defer repo.db.Close()

	var oneData T
	if err := repo.db.DB.First(&oneData, id); err != nil {
		return oneData, err.Error
	}

	return oneData, nil
}

func (repo *Repository[T]) Delete(id int) (int, error) {
	data, err := repo.ReadOne(id)
	if err != nil {
		return 0, err
	}

	repo.db.Connect()
	defer repo.db.Close()

	result := repo.db.DB.Delete(data)

	return int(result.RowsAffected), result.Error
}
