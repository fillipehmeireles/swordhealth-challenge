package controllers

import (
	dtos "SwordHealth/pkg/dtos/task"
	"SwordHealth/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskControllers struct {
	service *services.TaskServices
}

func NewTaskControllers() *TaskControllers {
	return &TaskControllers{service: services.NewTaskServices()}
}

func (controller *TaskControllers) ReadAll(c *gin.Context) {
	tasks, _ := controller.service.GetAll()
	c.JSON(http.StatusOK, tasks)
}

func (controller *TaskControllers) ReadOne(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	task, err := controller.service.GetOne(idParsed)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (controller *TaskControllers) Create(c *gin.Context) {
	var task dtos.CreateTaskDto

	task.TechnicianID = c.GetInt("technician_id")
	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := controller.service.Create(task)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Lines Affected": result})
}

func (controller *TaskControllers) Update(c *gin.Context) {

	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	var task dtos.UpdateTaskDto

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	task.TechnicianID = c.GetInt("technician_id")
	result, err := controller.service.Update(idParsed, task)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Lines Affected": result})

}

func (controller *TaskControllers) Delete(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	result, err := controller.service.Delete(idParsed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Lines Affected": result})

}

func (controller *TaskControllers) ReadAllTasksOfTechnician(c *gin.Context) {
	technicianID := c.GetInt("technician_id")
	tasks, _ := controller.service.GetAllTasksOfTechnician(technicianID)

	c.JSON(http.StatusOK, tasks)
}

func (controller *TaskControllers) ReadOneTaskOfTechnician(c *gin.Context) {
	technicianID := c.GetInt("technician_id")
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	task, err := controller.service.GetOneTaskOfTechnician(idParsed, technicianID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (controller *TaskControllers) ChangeTaskStatus(c *gin.Context) {

	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	var task dtos.TaskStatusDto

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	techID := c.GetInt("technician_id")
	result, err := controller.service.ChangeTaskStatus(idParsed, techID, task)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Lines Affected": result})

}
