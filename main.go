package main

import (
	"GoGinToDoList/cmd"
	"GoGinToDoList/config"
	"GoGinToDoList/controller"
	"GoGinToDoList/middlewares"
	"GoGinToDoList/repository"
	"GoGinToDoList/routes"
	"GoGinToDoList/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDBConnection()
	defer config.CloseDBConnection(db)

	if len(os.Args) > 1 {
		cmd.Command(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()
		// Implement Dependency injection
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    service.UserService       = service.NewUserService(userRepository, jwtService)
		userController controller.UserController = controller.NewUserController(userService)
	)

	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	routes.User(server, userController, jwtService)

	server.Static("/assets", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
