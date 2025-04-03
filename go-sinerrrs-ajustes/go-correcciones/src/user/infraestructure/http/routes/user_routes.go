package routes

import (
	"api/src/user/infraestructure/http/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, controller *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", controller.Create)
		userRoutes.GET("/", controller.GetAll)
		userRoutes.GET("/:id", controller.GetByID)
		userRoutes.PUT("/:id", controller.Update)
		userRoutes.DELETE("/:id", controller.Delete)
	}
}
