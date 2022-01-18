package controller

import (
	"fmt"
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
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type userController struct {
	userService service.UserService
	// jwtService  service.JWTService
}

//NewUserController create a new instances of UserController
func NewUserController(userServ service.UserService) UserController {
	return &userController{
		userService: userServ,
		// jwtService:  jwtServ,
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

func (c *userController) Insert(context *gin.Context) {
	var userCreateDTO dto.UserCreateDTO

	// fill variablle userCreateDTO
	errDTO := context.ShouldBind(&userCreateDTO)
	
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		fmt.Println("res =", res)
		context.JSON(http.StatusBadRequest, res)
	} else {
		if !c.userService.IsDuplicateEmail(userCreateDTO.Email) {
			response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
			context.JSON(http.StatusConflict, response)
		} else {
			// Validate password should 1 Uppercase, 2 Numbers and 1 Symbol
			// password := userCreateDTO.Password
			// var regex, _ = regexp.Compile(``)
	
			result := c.userService.Insert(userCreateDTO)
			// token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
			// createdUser.Token = token
			response := helper.BuildResponse(true, "OK!", result)
			context.JSON(http.StatusCreated, response)
		}
	}
	
	

		


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

	// Take id from url parameter and convert the data type from string to int
	intParam, err_id := strconv.ParseInt(context.Param("id"), 0, 64)
	if err_id == nil {
		userUpdateDTO.ID = intParam
	}

	// fmt.Println(userUpdateDTO.ID)
	// fmt.Println(userUpdateDTO.First_name)
	// fmt.Println(userUpdateDTO.Last_name)
	// fmt.Println(userUpdateDTO.Password)
	// fmt.Println(userUpdateDTO.Type)
	// fmt.Println(userUpdateDTO.Username)
	// fmt.Println(userUpdateDTO.Id_number)
	// os.Exit(1)

	result := c.userService.Update(userUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)

}