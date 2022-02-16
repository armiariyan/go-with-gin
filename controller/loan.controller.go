package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/helper"
	"gitlab.com/armiariyan/intern_golang/service"
)

//LoanController interface is a contract what this controller can do
type LoanController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type loanController struct {
	loanService service.LoanService
	jwtService  service.JWTService
}

//NewLoanController creates a new instance of LoanController
func NewLoanController(loanService service.LoanService, jwtService service.JWTService) LoanController {
	return &loanController{
		loanService: loanService,
		jwtService:  jwtService,
	}
}


func (c *loanController) Insert(context *gin.Context) {
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")

	// Validate token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	var loanDTO dto.LoanCreateDTO
	
	// Fill registerDTO variable
	errDTO := context.ShouldBind(&loanDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	
	// Validate confirmed amount shouldnt higher than loan_amount
	if loanDTO.Confirmed_amount > loanDTO.Loan_amount {
		response := helper.BuildErrorResponse("Failed to process request", "confirmed amount shouldnt higher than loan_amount", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	createdLoan := c.loanService.CreateLoan(loanDTO)

	response := helper.BuildResponse(true, "OK!", createdLoan)
	context.JSON(http.StatusCreated, response)
}

func (c *loanController) All(context *gin.Context) {
	var loans []entity.Loan = c.loanService.All()
	res := helper.BuildResponse(true, "OK", loans)
	context.JSON(http.StatusOK, res)
}

func (c *loanController) Update(context *gin.Context) {
	// Declare variable
	var loanUpdateDTO dto.LoanUpdateDTO
	
	
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
	errDTO := context.ShouldBind(&loanUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	// Take id from url parameter and convert the data type from string to int
	id := context.Param("id")
	
	// Validate if data exist
	result_checkId := c.loanService.FindByID(id)
	if (result_checkId == entity.Loan{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}

	loanUpdateDTO.Id_loan = id

	result := c.loanService.Update(loanUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)


}

func (c *loanController) Delete(context *gin.Context) {
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
	result_checkId := c.loanService.FindByID(id)

	if (result_checkId == entity.Loan{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	result_checkId.Id_loan = id
	c.loanService.Delete(result_checkId)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	
}


