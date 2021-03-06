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
	db             			*gorm.DB                  	  	  = config.SetupDatabaseConnection()
	userRepository 			repository.UserRepository 	      = repository.NewUserRepository(db)
	borrowerRepository	  	repository.BorrowerRepository 	  = repository.NewBorrowerRepository(db)
	lenderRepository  		repository.LenderRepository	  	  = repository.NewLenderRepository(db)
	requestLoanRepository  	repository.RequestLoanRepository  = repository.NewRequestLoanRepository(db)
	transactionRepository  	repository.TransactionRepository  = repository.NewTransactionRepository(db)
	loanPaymentRepository  	repository.Loan_paymentRepository = repository.NewLoan_paymentRepository(db)
	loanRepository 		 	repository.LoanRepository	  	  = repository.NewLoanRepository(db)

	jwtService     		service.JWTService        	  = service.NewJWTService()
	userService    		service.UserService       	  = service.NewUserService(userRepository)
	authService    		service.AuthService       	  = service.NewAuthService(userRepository)
	borrowerService    	service.BorrowerService       = service.NewBorrowerService(borrowerRepository)
	lenderService    	service.LenderService         = service.NewLenderService(lenderRepository)
	requestLoanService  service.RequestLoanService    = service.NewRequestLoanService(requestLoanRepository)
	transactionService  service.TransactionService    = service.NewTransactionService(transactionRepository)
	loanPaymentService  service.Loan_paymentService   = service.NewLoan_paymentService(loanPaymentRepository)
	loanService		    service.LoanService 		  = service.NewLoanService(loanRepository)

	userController		  controller.UserController 	    = controller.NewUserController(userService, jwtService)
	authController 		  controller.AuthController 	    = controller.NewAuthController(authService, jwtService)
	borrowerController 	  controller.BorrowerController		= controller.NewBorrowerController(borrowerService, jwtService)
	lenderController 	  controller.LenderController       = controller.NewLenderController(lenderService, jwtService)
	requestLoanController controller.RequestLoanController  = controller.NewRequestLoanController(requestLoanService, jwtService)
	transactionController controller.TransactionController  = controller.NewTransactionController(transactionService, jwtService)
	loanPaymentController controller.Loan_paymentController = controller.NewLoan_paymentController(loanPaymentService, jwtService)
	loanController		  controller.LoanController 		= controller.NewLoanController(loanService, jwtService)
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
		borrowerRoutes.DELETE("/:id", borrowerController.Delete)
	}

	lenderRoutes := r.Group("api/lender")
	{
		lenderRoutes.GET("/", lenderController.All)
		lenderRoutes.POST("/", lenderController.Insert)
		lenderRoutes.PUT("/:id", lenderController.Update)
		lenderRoutes.DELETE("/:id", lenderController.Delete)
	}
	
	requestLoanRoutes := r.Group("api/request-loan")
	{
		requestLoanRoutes.GET("/", requestLoanController.All)
		requestLoanRoutes.POST("/", requestLoanController.Insert)
		requestLoanRoutes.PUT("/:id", requestLoanController.Update)
		requestLoanRoutes.DELETE("/:id", requestLoanController.Delete)
	}

	transactionRoutes := r.Group("api/transaction")
	{
		transactionRoutes.GET("/", transactionController.All)
		transactionRoutes.POST("/", transactionController.Insert)
		transactionRoutes.PUT("/:id", transactionController.Update)
		transactionRoutes.DELETE("/:id", transactionController.Delete)
	}

	loanPaymentRoutes := r.Group("api/loan-payment")
	{
		loanPaymentRoutes.GET("/", loanPaymentController.All)
		loanPaymentRoutes.POST("/", loanPaymentController.Insert)
		loanPaymentRoutes.PUT("/:id", loanPaymentController.Update)
		loanPaymentRoutes.DELETE("/:id", loanPaymentController.Delete)
	}

	loanRoutes := r.Group("api/loan")
	{
		loanRoutes.GET("/", loanController.All)
		loanRoutes.POST("/", loanController.Insert)
		loanRoutes.PUT("/:id", loanController.Update)
		loanRoutes.DELETE("/:id", loanController.Delete)
	}

	r.Run()
}