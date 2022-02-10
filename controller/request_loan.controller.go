package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	token, errToken := c.jwtService.ValidateToken(authHeader)
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

	// Make the string input Capitallize
	requestLoanDTO.Sumber_dana =  strings.Title(requestLoanDTO.Sumber_dana)

	// Validate Sumber Dana value
	if requestLoanDTO.Sumber_dana != "Dalam Negeri" && requestLoanDTO.Sumber_dana != "Luar Negeri" {
		response := helper.BuildErrorResponse("Failed to process request", "Sumber_dana should be Dalam Negeri or Luar Negeri", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Get ID User from Token JWT
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("claims=",claims)
	
	id_user, err := strconv.ParseInt(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	// Assign ID user
	requestLoanDTO.Id_user = id_user
	

	// Create Id_RequestLoan
	id_requestLoan := c.requestLoanService.CreateIdRequestLoan()
	requestLoanDTO.Id_requestLoan = id_requestLoan
	
	// Insert data
	createdRequestLoan := c.requestLoanService.CreateRequestLoan(requestLoanDTO)

	response := helper.BuildResponse(true, "OK!", createdRequestLoan)
	context.JSON(http.StatusCreated, response)


}

func (c *requestLoanController) All(context *gin.Context) {
	var requestLoans []entity.RequestLoan = c.requestLoanService.All()
	res := helper.BuildResponse(true, "OK", requestLoans)
	context.JSON(http.StatusOK, res)
}



func (c *requestLoanController) Update(context *gin.Context) {
	// Declare variable
	var requestLoanUpdateDTO dto.RequestLoanUpdateDTO
	
	
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")
	
	// Validate token
	token, errToken := c.jwtService.ValidateToken(authHeader)
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
	// Make the string input Capitallize
	requestLoanUpdateDTO.Sumber_dana =  strings.Title(requestLoanUpdateDTO.Sumber_dana)

	// Validate Sumber Dana value
	if requestLoanUpdateDTO.Sumber_dana != "Dalam Negeri" && requestLoanUpdateDTO.Sumber_dana != "Luar Negeri" {
		response := helper.BuildErrorResponse("Failed to process request", "Sumber_dana should be Dalam Negeri or Luar Negeri", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Take id from url parameter and convert the data type from string to int
	id := context.Param("id")
	

	// Validate if data exist
	result_checkId := c.requestLoanService.FindByID(id)

	if (result_checkId == entity.RequestLoan{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	requestLoanUpdateDTO.Id_requestLoan = id


	// Get ID User from Token JWT
	claims := token.Claims.(jwt.MapClaims)
	// fmt.Println("claims=",claims)
	
	id_user, err := strconv.ParseInt(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	// Assign ID user
	requestLoanUpdateDTO.Id_user = id_user

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

	if (result_checkId == entity.RequestLoan{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	result_checkId.Id_requestLoan = id
	c.requestLoanService.Delete(result_checkId)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	
}


