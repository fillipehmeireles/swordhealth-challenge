package repository

import (
	"SwordHealth/pkg/api/drivers"
	dtos "SwordHealth/pkg/api/dtos/task"
	"SwordHealth/pkg/api/models"
	"errors"
	"time"
)

type TaskRepository struct {
	Repository[models.Task]
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{Repository[models.Task]{db: drivers.NewDBDriver()}}
}

func (repo *TaskRepository) Create(taskDto dtos.CreateTaskDto) (int, error) {
	repo.db.Connect()
	defer repo.db.Close()
	task := models.Task{
		Title:        taskDto.Title,
		Summary:      taskDto.Summary,
		TechnicianID: taskDto.TechnicianID,
	}

	result := repo.db.DB.Create(&task)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *TaskRepository) Update(id int, taskDto dtos.UpdateTaskDto) (int, error) {
	task, err := repo.ReadOne(id)
	if err != nil {
		return 0, err
	}

	repo.db.Connect()
	defer repo.db.Close()

	if task.TechnicianID != taskDto.TechnicianID {
		return 0, errors.New("Could not update a task which is not assigned to you.")
	}

	taskModel := models.Task{
		Title:        taskDto.Title,
		Summary:      taskDto.Summary,
		TechnicianID: taskDto.TechnicianID,
	}

	result := repo.db.DB.Model(&task).Updates(taskModel)

	return int(result.RowsAffected), result.Error

}

func (repo *TaskRepository) ReadAll() ([]models.Task, error) {
	repo.db.Connect()
	defer repo.db.Close()

	var allData []models.Task
	repo.db.DB.Preload("Technician").Find(&allData)
	return allData, nil
}

func (repo *TaskRepository) ReadOne(id int) (models.Task, error) {
	repo.db.Connect()
	defer repo.db.Close()

	var oneData models.Task
	if err := repo.db.DB.Preload("Technician").First(&oneData, id); err != nil {
		return oneData, err.Error
	}

	return oneData, nil
}

func (repo *TaskRepository) ReadAllTasksOfTechnician(technicianID int) ([]models.Task, error) {
	repo.db.Connect()
	defer repo.db.Close()

	var allData []models.Task
	repo.db.DB.Preload("Technician").Where("technician_id = ?", technicianID).Find(&allData)
	return allData, nil
}

func (repo *TaskRepository) ReadOneTaskOfTechnician(id, technicianID int) (models.Task, error) {
	repo.db.Connect()
	defer repo.db.Close()

	var oneData models.Task
	if err := repo.db.DB.Preload("Technician").Where("technician_id = ?", technicianID).First(&oneData, id); err != nil {
		return oneData, err.Error
	}
	return oneData, nil
}

func (repo *TaskRepository) ChangeTaskStatus(id int, taskDto dtos.TaskStatusDto, techID int) (models.Task, int, error) {
	task, err := repo.ReadOne(id)
	if err != nil {
		return models.Task{}, 0, err
	}
	if task.TechnicianID != techID {
		return models.Task{}, 0, errors.New("Could not update a task which is not assigned to you.")
	}

	repo.db.Connect()
	defer repo.db.Close()

	taskModel := models.Task{
		Status:      taskDto.Status,
		PerformedAt: time.Now(),
	}

	result := repo.db.DB.Model(&task).Updates(taskModel)

	return task, int(result.RowsAffected), result.Error

}
