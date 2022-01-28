package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/config"
	"gitlab.com/armiariyan/intern_golang/controller"
	"gitlab.com/armiariyan/intern_golang/repository"
	"gitlab.com/armiariyan/intern_golang/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)


	userController controller.UserController = controller.NewUserController(userService, jwtService)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group(("api/auth"))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user")
	{
		userRoutes.GET("/", userController.All)
		userRoutes.PUT("/:id", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
	}

	

	r.Run()
}