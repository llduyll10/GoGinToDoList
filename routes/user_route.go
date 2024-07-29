package routes

import (
	"GoGinToDoList/controller"
	"GoGinToDoList/middlewares"
	"GoGinToDoList/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
		routes.GET("/me", middlewares.Authenticate(jwtService), userController.Me)
	}
}
