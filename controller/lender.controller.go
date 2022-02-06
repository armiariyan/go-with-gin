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

//LenderController interface is a contract what this controller can do
type LenderController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type lenderController struct {
	lenderService service.LenderService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewLenderController(lenderService service.LenderService, jwtService service.JWTService) LenderController {
	return &lenderController{
		lenderService: lenderService,
		jwtService:  jwtService,
	}
}


func (c *lenderController) Insert(context *gin.Context) {
	// Take Token from Header named Authorization
	authHeader := context.GetHeader("Authorization")

	// Validate token
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	var lenderDTO dto.LenderCreateDTO

	// Fill registerDTO variable
	errDTO := context.ShouldBind(&lenderDTO)
	if errDTO != nil {
	response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
	context.AbortWithStatusJSON(http.StatusBadRequest, response)
	return
	}

	// Make the string input Capitallize
	lenderDTO.Sumber_dana =  strings.Title(lenderDTO.Sumber_dana)

	// Validate Sumber Dana value
	if lenderDTO.Sumber_dana != "Dalam Negeri" && lenderDTO.Sumber_dana != "Luar Negeri" {
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
	lenderDTO.Id_user = id_user
	

	// Create Id_Lender
	id_lender := c.lenderService.CreateIdLender()
	lenderDTO.Id_lender = id_lender
	
	// Insert data
	createdLender := c.lenderService.CreateLender(lenderDTO)

	response := helper.BuildResponse(true, "OK!", createdLender)
	context.JSON(http.StatusCreated, response)


}

func (c *lenderController) All(context *gin.Context) {
	var lenders []entity.Lender = c.lenderService.All()
	res := helper.BuildResponse(true, "OK", lenders)
	context.JSON(http.StatusOK, res)
}



func (c *lenderController) Update(context *gin.Context) {
	// Declare variable
	var lenderUpdateDTO dto.LenderUpdateDTO
	
	
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
	errDTO := context.ShouldBind(&lenderUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	// Make the string input Capitallize
	lenderUpdateDTO.Sumber_dana =  strings.Title(lenderUpdateDTO.Sumber_dana)

	// Validate Sumber Dana value
	if lenderUpdateDTO.Sumber_dana != "Dalam Negeri" && lenderUpdateDTO.Sumber_dana != "Luar Negeri" {
		response := helper.BuildErrorResponse("Failed to process request", "Sumber_dana should be Dalam Negeri or Luar Negeri", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Take id from url parameter and convert the data type from string to int
	id := context.Param("id")
	

	// Validate if data exist
	result_checkId := c.lenderService.FindByID(id)

	if (result_checkId == entity.Lender{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	lenderUpdateDTO.Id_lender = id


	// Get ID User from Token JWT
	claims := token.Claims.(jwt.MapClaims)
	// fmt.Println("claims=",claims)
	
	id_user, err := strconv.ParseInt(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	// Assign ID user
	lenderUpdateDTO.Id_user = id_user

	result := c.lenderService.Update(lenderUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *lenderController) Delete(context *gin.Context) {
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
	result_checkId := c.lenderService.FindByID(id)

	if (result_checkId == entity.Lender{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	result_checkId.Id_lender = id
	c.lenderService.Delete(result_checkId)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	
}


