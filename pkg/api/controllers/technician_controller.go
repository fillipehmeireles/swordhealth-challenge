package controllers

import (
	dtos "SwordHealth/pkg/api/dtos/technician"
	"SwordHealth/pkg/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TechnicianControllers struct {
	service *services.TechnicianServices
}

func NewTechnicianControllers() *TechnicianControllers {
	return &TechnicianControllers{service: services.NewTechnicianServices()}
}

func (controller *TechnicianControllers) ReadAll(c *gin.Context) {
	technicians, _ := controller.service.GetAll()
	c.JSON(http.StatusOK, technicians)
}

func (controller *TechnicianControllers) ReadOne(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	technician, err := controller.service.GetOne(idParsed)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, technician)
}

func (controller *TechnicianControllers) Create(c *gin.Context) {
	var technician dtos.CreateTechnicianDto

	if err := c.BindJSON(&technician); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := controller.service.Create(technician)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Lines Affected": result})
}

func (controller *TechnicianControllers) Update(c *gin.Context) {
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse param id to int"})
		return
	}

	var technician dtos.UpdateTechnicianDto

	if err := c.BindJSON(&technician); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := controller.service.Update(idParsed, technician)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Lines Affected": result})

}

func (controller *TechnicianControllers) Delete(c *gin.Context) {
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
