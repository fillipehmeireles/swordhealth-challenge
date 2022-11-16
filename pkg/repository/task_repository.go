package repository

import (
	"SwordHealth/pkg/drivers"
	dtos "SwordHealth/pkg/dtos/task"
	"SwordHealth/pkg/models"
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
