package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/config"
	"gitlab.com/armiariyan/intern_golang/controller"
	"gitlab.com/armiariyan/intern_golang/repository"
	"gitlab.com/armiariyan/intern_golang/seeder"
	"gitlab.com/armiariyan/intern_golang/service"
	"gorm.io/gorm"
)

var (
	db             		*gorm.DB                  	  = config.SetupDatabaseConnection()
	userRepository 		repository.UserRepository 	  = repository.NewUserRepository(db)
	borrowerRepository  repository.BorrowerRepository = repository.NewBorrowerRepository(db)
	lenderRepository  	repository.LenderRepository	  = repository.NewLenderRepository(db)
	
	jwtService     		service.JWTService        	  = service.NewJWTService()
	userService    		service.UserService       	  = service.NewUserService(userRepository)
	authService    		service.AuthService       	  = service.NewAuthService(userRepository)
	borrowerService    	service.BorrowerService       = service.NewBorrowerService(borrowerRepository)
	lenderService    	service.LenderService         = service.NewLenderService(lenderRepository)


	userController controller.UserController = controller.NewUserController(userService, jwtService)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	borrowerController controller.BorrowerController = controller.NewBorrowerController(borrowerService, jwtService)
	lenderController controller.LenderController = controller.NewLenderController(lenderService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	seeder.DBSeed(db)
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

	borrowerRoutes := r.Group("api/borrower")
	{
		borrowerRoutes.GET("/", borrowerController.All)
		borrowerRoutes.POST("/", borrowerController.Insert)
		borrowerRoutes.PUT("/:id", borrowerController.Update)
		// borrowerRoutes.DELETE("/:id", borrowerController.Delete)
	}

	lenderRoutes := r.Group("api/lender")
	{
		lenderRoutes.POST("/", lenderController.Insert)
	}
	

	r.Run()
}