package services

import (
	"SwordHealth/pkg/repository"
)

type Services[T any, U repository.IRepository[T]] struct {
	repository U
}

func (service *Services[T, U]) GetAll() ([]T, error) {
	allData, _ := service.repository.ReadAll()

	return allData, nil
}

func (service *Services[T, U]) GetOne(id int) (T, error) {
	return service.repository.ReadOne(id)
}

func (service *Services[T, U]) Delete(id int) (int, error) {
	return service.repository.Delete(id)
}
