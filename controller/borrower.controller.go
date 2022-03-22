package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/helper"
	"gitlab.com/armiariyan/intern_golang/service"
)

//BorrowerController interface is a contract what this controller can do
type BorrowerController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type borrowerController struct {
	borrowerService service.BorrowerService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewBorrowerController(borrowerService service.BorrowerService, jwtService service.JWTService) BorrowerController {
	return &borrowerController{
		borrowerService: borrowerService,
		jwtService:  jwtService,
	}
}


func (c *borrowerController) Insert(context *gin.Context) {
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")

	// Validate token
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	var borrowerDTO dto.BorrowerCreateDTO

	// Fill registerDTO variable
	errDTO := context.ShouldBind(&borrowerDTO)
	if errDTO != nil {
	response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
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
	borrowerDTO.Id_user = id_user
	

	// Create Id_Borrower
	id_borrower := c.borrowerService.CreateIdBorrower()
	borrowerDTO.Id_borrower = id_borrower
	
	// Insert data
	createdBorrower := c.borrowerService.CreateBorrower(borrowerDTO)

	response := helper.BuildResponse(true, "OK!", createdBorrower)
	context.JSON(http.StatusCreated, response)


}

func (c *borrowerController) All(context *gin.Context) {
	var borrowers []entity.Borrower = c.borrowerService.All()
	res := helper.BuildResponse(true, "OK", borrowers)
	context.JSON(http.StatusOK, res)
}



func (c *borrowerController) Update(context *gin.Context) {
	// Declare variable
	var borrowerUpdateDTO dto.BorrowerUpdateDTO
	
	// Fill the variable
	errDTO := context.ShouldBind(&borrowerUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")

	// Validate token
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	// Take id from url parameter and convert the data type from string to int
	id := context.Param("id")
	

	// Validate if data exist
	result_checkId := c.borrowerService.FindByID(id)

	if (result_checkId == entity.Borrower{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	borrowerUpdateDTO.Id_borrower = id


	// Get ID User from Token JWT
	claims := token.Claims.(jwt.MapClaims)
	// fmt.Println("claims=",claims)
	
	id_user, err := strconv.ParseInt(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	// Assign ID user
	borrowerUpdateDTO.Id_user = id_user

	result := c.borrowerService.Update(borrowerUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *borrowerController) Delete(context *gin.Context) {
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
	result_checkId := c.borrowerService.FindByID(id)

	if (result_checkId == entity.Borrower{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	result_checkId.Id_borrower = id
	c.borrowerService.Delete(result_checkId)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	
}


