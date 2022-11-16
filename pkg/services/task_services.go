package services

import (
	dtos "SwordHealth/pkg/dtos/task"
	"SwordHealth/pkg/models"
	"SwordHealth/pkg/repository"
)

type TaskServices struct {
	Services[models.Task, *repository.TaskRepository]
}

func NewTaskServices() *TaskServices {
	return &TaskServices{Services[models.Task, *repository.TaskRepository]{repository: repository.NewTaskRepository()}}
}

func (service *TaskServices) Create(taskDto dtos.CreateTaskDto) (int, error) {
	technicianServices := NewTechnicianServices()
	_, err := technicianServices.GetOne(taskDto.TechnicianID)
	if err != nil {
		return 0, err
	}

	return service.repository.Create(taskDto)
}

func (service *TaskServices) Update(id int, taskDto dtos.UpdateTaskDto) (int, error) {
	return service.repository.Update(id, taskDto)
}
