package routes

import (
	"SwordHealth/pkg/api/controllers"
	"SwordHealth/pkg/api/middlewares"

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
	r.GET("/", controllers.IndexController)

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
		taskRoutes.POST(createRoute, middlewares.TechAuth, taskController.Create)
		taskRoutes.GET(readAllRoute, taskController.ReadAll)
		taskRoutes.GET(readOneRoute, taskController.ReadOne)
		taskRoutes.PUT(updateRoute, middlewares.TechAuth, taskController.Update)
		taskRoutes.DELETE(deleteRoute, taskController.Delete)
		taskRoutes.GET("/technician", middlewares.TechAuth, taskController.ReadAllTasksOfTechnician)
		taskRoutes.GET("/technician/:id", middlewares.TechAuth, taskController.ReadOneTaskOfTechnician)
		taskRoutes.PATCH("/change_status/:id", middlewares.TechAuth, taskController.ChangeTaskStatus)
	}

}
