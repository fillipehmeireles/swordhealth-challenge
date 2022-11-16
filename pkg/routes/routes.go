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

	technicianRoutes := r.Group("/technician")
	{
		technicianController := controllers.NewTechnicianControllers()
		technicianRoutes.POST(createRoute, technicianController.Create)
		technicianRoutes.GET(readAllRoute, technicianController.ReadAll)
		technicianRoutes.GET(readOneRoute, technicianController.ReadOne)
		technicianRoutes.PUT(updateRoute, technicianController.Update)
		technicianRoutes.DELETE(deleteRoute, technicianController.Delete)
	}

	taskRoutes := r.Group("/task")
	{
		taskController := controllers.NewTaskControllers()
		taskRoutes.POST(createRoute, taskController.Create)
		taskRoutes.GET(readAllRoute, taskController.ReadAll)
		taskRoutes.GET(readOneRoute, taskController.ReadOne)
		taskRoutes.PUT(updateRoute, taskController.Update)
		taskRoutes.DELETE(deleteRoute, taskController.Delete)
	}

}
