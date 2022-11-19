package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"SwordHealth Challenge API": "v1.0"})
}
