package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/helper"
	"gitlab.com/armiariyan/intern_golang/service"
)

//UserController is a ...
type UserController interface {
	All(context *gin.Context)
	// FindByID(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserController create a new instances of UserController
func NewUserController(userServ service.UserService, jwtServ service.JWTService) UserController {
	return &userController{
		userService: userServ,
		jwtService:  jwtServ,
	}
}

func (c *userController) All(context *gin.Context) {
	var users []entity.User = c.userService.All()
	res := helper.BuildResponse(true, "OK", users)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Delete(context *gin.Context) {
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
	id, err_id := strconv.ParseInt(context.Param("id"), 0, 64)
	if err_id != nil {
		response := helper.BuildErrorResponse("ID Error!", err_id.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	// Validate if data exist and declare variable with entity type, because the delete service parameter is entity type
	var user entity.User = c.userService.FindByID(id)
	if (user == entity.User{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	user.ID = id
	c.userService.Delete(user)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	
}


func (c *userController) Update(context *gin.Context) {
	// Declare variable
	var userUpdateDTO dto.UserUpdateDTO
	
	// Fill the variable
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

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
	id, err_id := strconv.ParseInt(context.Param("id"), 0, 64)
	if err_id != nil {
		response := helper.BuildErrorResponse("ID Error!", err_id.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	// Validate if data exist
	result_checkId := c.userService.FindByID(id)
	if (result_checkId == entity.User{}) {
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
		return
	}
	
	userUpdateDTO.ID = id
	result := c.userService.Update(userUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
	

}