package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TechAuth(c *gin.Context) {
	technicianHeader := c.Request.Header["Technician_id"]
	if len(technicianHeader) == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Please provide the technician ID"})
		return
	}

	technicianID := technicianHeader[0]
	techIDParsed, err := strconv.Atoi(technicianID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse header ID to INT"})
		return
	}

	c.Set("technician_id", techIDParsed)

	c.Next()
}
