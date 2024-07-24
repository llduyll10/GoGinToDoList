package main

import (
	"GoGinToDoList/config"
	"GoGinToDoList/controller"
	"GoGinToDoList/middlewares"
	"GoGinToDoList/repository"
	"GoGinToDoList/routes"
	"GoGinToDoList/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	db := config.SetUpDBConnection()
	defer config.CloseDBConnection(db)

	var (
		// Implement Dependency injection
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    service.UserService       = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService)
	)

	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	routes.User(server, userController)

	server.Static("/assets", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
