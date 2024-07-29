package routes

import (
	"GoGinToDoList/controller"
	"GoGinToDoList/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
	}
}
