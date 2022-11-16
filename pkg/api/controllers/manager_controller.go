package controllers

import (
	dtos "SwordHealth/pkg/api/dtos/manager"
	"SwordHealth/pkg/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ManagerControllers struct {
	service *services.ManagerServices
}

func NewManagerControllers() *ManagerControllers {
	return &ManagerControllers{service: services.NewManagerServices()}
}

func (controller *ManagerControllers) ReadAll(c *gin.Context) {
	managers, _ := controller.service.GetAll()
	c.JSON(http.StatusOK, managers)
}

func (controller *ManagerControllers) ReadOne(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	manager, err := controller.service.GetOne(idParsed)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, manager)
}

func (controller *ManagerControllers) Create(c *gin.Context) {
	var manager dtos.CreateManagerDTO

	if err := c.BindJSON(&manager); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := controller.service.Create(manager)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Lines Affected": result})
}

func (controller *ManagerControllers) Update(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	var manager dtos.UpdateManagerDTO

	if err := c.BindJSON(&manager); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := controller.service.Update(idParsed, manager)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Lines Affected": result})

}

func (controller *ManagerControllers) Delete(c *gin.Context) {
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
