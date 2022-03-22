package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/helper"
	"gitlab.com/armiariyan/intern_golang/service"
)

//Loan_paymentController interface is a contract what this controller can do
type Loan_paymentController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type loan_paymentController struct {
	loan_paymentService service.Loan_paymentService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewLoan_paymentController(loan_paymentService service.Loan_paymentService, jwtService service.JWTService) Loan_paymentController {
	return &loan_paymentController{
		loan_paymentService: loan_paymentService,
		jwtService:  jwtService,
	}
}


func (c *loan_paymentController) Insert(context *gin.Context) {
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")

	// Validate token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	var loan_paymentDTO dto.Loan_paymentCreateDTO
	
	// Fill registerDTO variable
	errDTO := context.ShouldBind(&loan_paymentDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create Id_Loan_payment
	id_loan_payment := c.loan_paymentService.CreateIdLoan_payment()
	loan_paymentDTO.Id_payment = id_loan_payment

	createdLoan_payment := c.loan_paymentService.CreateLoan_payment(loan_paymentDTO)

	response := helper.BuildResponse(true, "OK!", createdLoan_payment)
	context.JSON(http.StatusCreated, response)
}

func (c *loan_paymentController) All(context *gin.Context) {
	var loan_payments []entity.Loan_payment = c.loan_paymentService.All()
	res := helper.BuildResponse(true, "OK", loan_payments)
	context.JSON(http.StatusOK, res)
}

func (c *loan_paymentController) Update(context *gin.Context) {
	// Declare variable
	var loan_paymentUpdateDTO dto.Loan_paymentUpdateDTO
	
	
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
	errDTO := context.ShouldBind(&loan_paymentUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	// Take id from url parameter and convert the data type from string to int
	id := context.Param("id")
	
	// Validate if data exist
	result_checkId := c.loan_paymentService.FindByID(id)
	if (result_checkId == entity.Loan_payment{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}

	loan_paymentUpdateDTO.Id_payment = id

	result := c.loan_paymentService.Update(loan_paymentUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)


}

func (c *loan_paymentController) Delete(context *gin.Context) {
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
	result_checkId := c.loan_paymentService.FindByID(id)

	if (result_checkId == entity.Loan_payment{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	result_checkId.Id_payment = id
	c.loan_paymentService.Delete(result_checkId)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	
}


