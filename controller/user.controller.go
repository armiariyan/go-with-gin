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
	// Konversi ?
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	
	// Error jika parameter id tidak ada
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found / Param id error", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
		return
	}

	// Cek apakah data dengan ID tersebut ada sekaligus declare
	var user entity.User = c.userService.FindByID(id)

	if (user == entity.User{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		user.ID = id
		// Delete user
		c.userService.Delete(user)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
	// authHeader := context.GetHeader("Authorization")
	// token, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// userID := fmt.Sprintf("%v", claims["user_id"])
	// if c.bookService.IsAllowedToEdit(userID, book.ID) {
	// 	c.bookService.Delete(book)
	// 	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	// 	context.JSON(http.StatusOK, res)
	// } else {
	// 	response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
	// 	context.JSON(http.StatusForbidden, response)
	// }
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

	// Validate token step
	// Take Header Named Authorization
	authHeader := context.GetHeader("Authorization")
	// Validate token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	// If there is error, show it
	if errToken != nil {
		response := helper.BuildErrorResponse("Token Error!", errToken.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}

	// claims := token.Claims.(jwt.MapClaims)
	// fmt.Println("token before claims=",claims)
	// fmt.Println("claims=",claims)

	// Take id from url parameter and convert the data type from string to int
	intParam, err_id := strconv.ParseInt(context.Param("id"), 0, 64)
	if err_id != nil {
		response := helper.BuildErrorResponse("ID Error!", err_id.Error(), helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
		return
	}
	
	userUpdateDTO.ID = intParam
	result := c.userService.Update(userUpdateDTO)
	if result.ID == 0 { 
		response := helper.BuildErrorResponse("Failed to proccess request", "Record with given ID not found", helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
	} else {
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
	

}