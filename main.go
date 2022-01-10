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
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)

)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	userRoutes := r.Group("api/user")
	{
		userRoutes.GET("/", userController.All)
		userRoutes.POST("/", userController.Insert)
		// userRoutes.PUT("/", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
	}

	r.Run()
}