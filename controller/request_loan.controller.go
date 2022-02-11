package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/helper"
	"gitlab.com/armiariyan/intern_golang/service"
)

//RequestLoanController interface is a contract what this controller can do
type RequestLoanController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type requestLoanController struct {
	requestLoanService service.RequestLoanService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewRequestLoanController(requestLoanService service.RequestLoanService, jwtService service.JWTService) RequestLoanController {
	return &requestLoanController{
		requestLoanService: requestLoanService,
		jwtService:  jwtService,
	}
}


func (c *requestLoanController) Insert(context *gin.Context) {
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")

	// Validate token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	var requestLoanDTO dto.RequestLoanCreateDTO

	// Fill registerDTO variable
	errDTO := context.ShouldBind(&requestLoanDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create Id_RequestLoan
	id_requestLoan := c.requestLoanService.CreateIdRequestLoan()
	requestLoanDTO.Id_loan = id_requestLoan

	// Validate loan_duration cannot be 0
	if requestLoanDTO.Loan_duration == 0 {
		response := helper.BuildErrorResponse("Failed to process request", "Loan_duration cannot be 0", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Validate loan_amount min 50.000
	if requestLoanDTO.Loan_amount < 50000 {
		response := helper.BuildErrorResponse("Failed to process request", "Loan_amount minimum is 50000", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	
	// Insert data
	createdRequestLoan := c.requestLoanService.CreateRequestLoan(requestLoanDTO)

	response := helper.BuildResponse(true, "OK!", createdRequestLoan)
	context.JSON(http.StatusCreated, response)


}

func (c *requestLoanController) All(context *gin.Context) {
	var requestLoans []entity.Request_loan = c.requestLoanService.All()
	res := helper.BuildResponse(true, "OK", requestLoans)
	context.JSON(http.StatusOK, res)
}



func (c *requestLoanController) Update(context *gin.Context) {
	// Declare variable
	var requestLoanUpdateDTO dto.RequestLoanUpdateDTO
	
	
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
	errDTO := context.ShouldBind(&requestLoanUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	// Take id from url parameter and convert the data type from string to int
	id := context.Param("id")
	
	// Validate if data exist
	result_checkId := c.requestLoanService.FindByID(id)
	if (result_checkId == entity.Request_loan{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	requestLoanUpdateDTO.Id_loan = id

	// Validate loan_duration cannot be 0
	if requestLoanUpdateDTO.Loan_duration == 0 {
		response := helper.BuildErrorResponse("Failed to process request", "Loan_duration cannot be 0", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Validate loan_amount min 50.000
	if requestLoanUpdateDTO.Loan_amount < 50000 {
		response := helper.BuildErrorResponse("Failed to process request", "Loan_amount minimum is 50000", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result := c.requestLoanService.Update(requestLoanUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *requestLoanController) Delete(context *gin.Context) {
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
	result_checkId := c.requestLoanService.FindByID(id)

	if (result_checkId == entity.Request_loan{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	result_checkId.Id_loan = id
	c.requestLoanService.Delete(result_checkId)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	
}


