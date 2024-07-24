package routes

import (
	"GoGinToDoList/controller"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController) {
	routes := route.Group("/api/user")
	{
		routes.POST("", userController.Register)
	}
}
