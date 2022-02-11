package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/helper"
	"gitlab.com/armiariyan/intern_golang/service"
)

//TransactionController interface is a contract what this controller can do
type TransactionController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewTransactionController(transactionService service.TransactionService, jwtService service.JWTService) TransactionController {
	return &transactionController{
		transactionService: transactionService,
		jwtService:  jwtService,
	}
}


func (c *transactionController) Insert(context *gin.Context) {
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")

	// Validate token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	var transactionDTO dto.TransactionCreateDTO

	// timez := time.Now().Format("2006-01-02")
	// transactionDTO.Borrower_deal_date = timez
	// transactionDTO.Due_date = timez
	
	// Fill registerDTO variable
	errDTO := context.ShouldBind(&transactionDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}


	// Create Id_Transaction
	id_transaction := c.transactionService.CreateIdTransaction()
	transactionDTO.Id_transaction = id_transaction

	// Validate interest rate
	if transactionDTO.Interest_rate > 1 {
		response := helper.BuildErrorResponse("Failed to process request", "Interest rate can not more than 1", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		// response := helper.BuildErrorResponse("Failed to process request", "below is the data", transactionDTO)
		// context.AbortWithStatusJSON(http.StatusBadRequest, response)
		// return
		// Insert data
		createdTransaction := c.transactionService.CreateTransaction(transactionDTO)

		response := helper.BuildResponse(true, "OK!", createdTransaction)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *transactionController) All(context *gin.Context) {
	var transactions []entity.Transaction = c.transactionService.All()
	res := helper.BuildResponse(true, "OK", transactions)
	context.JSON(http.StatusOK, res)
}

func (c *transactionController) Update(context *gin.Context) {
	// Declare variable
	var transactionUpdateDTO dto.TransactionUpdateDTO
	
	
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")
	
	// Validate token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}
	
	// Fill the variable
	errDTO := context.ShouldBind(&transactionUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	// Take id from url parameter and convert the data type from string to int
	id := context.Param("id")
	
	// Validate if data exist
	result_checkId := c.transactionService.FindByID(id)
	if (result_checkId == entity.Transaction{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	transactionUpdateDTO.Id_transaction = id

	// Validate interest rate
	if transactionUpdateDTO.Interest_rate > 1 {
		response := helper.BuildErrorResponse("Failed to process request", "Interest rate can not more than 1", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		// response := helper.BuildErrorResponse("Failed to process request", "below is the data", transactionDTO)
		// context.AbortWithStatusJSON(http.StatusBadRequest, response)
		// return

		// Update data
		result := c.transactionService.Update(transactionUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}


}

func (c *transactionController) Delete(context *gin.Context) {
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")
	
	// Validate token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	// Take id from url parameter and convert the data type from string to int
	id := context.Param("id")

	// Validate if data exist and declare variable with entity type, because the delete service parameter is entity type
	result_checkId := c.transactionService.FindByID(id)

	if (result_checkId == entity.Transaction{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	result_checkId.Id_transaction = id
	c.transactionService.Delete(result_checkId)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	
}


