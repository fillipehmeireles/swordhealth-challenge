package routes

import (
	"SwordHealth/pkg/controllers"

	"github.com/gin-gonic/gin"
)

const (
	readAllRoute string = "/read-all"
	readOneRoute string = "/read-one/:id"
	createRoute  string = "/create"
	updateRoute  string = "/update/:id"
	deleteRoute  string = "/delete/:id"
)

func Routes(r *gin.Engine) {
	managerRoutes := r.Group("/manager")
	{
		managerController := controllers.NewManagerControllers()
		managerRoutes.POST(createRoute, managerController.Create)
		managerRoutes.GET(readAllRoute, managerController.ReadAll)
		managerRoutes.GET(readOneRoute, managerController.ReadOne)
		managerRoutes.PUT(updateRoute, managerController.Update)
		managerRoutes.DELETE(deleteRoute, managerController.Delete)
	}
}
