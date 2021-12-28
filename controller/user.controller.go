package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/helper"
	"gitlab.com/armiariyan/intern_golang/service"
)

//UserController is a ...
type UserController interface {
	All(context *gin.Context)
	// FindByID(context *gin.Context)
	// Insert(context *gin.Context)
	// Update(context *gin.Context)
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

// func (c *bookController) FindByID(context *gin.Context) {
// 	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
// 	if err != nil {
// 		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
// 		context.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	var book entity.Book = c.bookService.FindByID(id)
// 	if (book == entity.Book{}) {
// 		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
// 		context.JSON(http.StatusNotFound, res)
// 	} else {
// 		res := helper.BuildResponse(true, "OK", book)
// 		context.JSON(http.StatusOK, res)
// 	}
// }

// func (c *bookController) Insert(context *gin.Context) {
// 	var bookCreateDTO dto.BookCreateDTO
// 	errDTO := context.ShouldBind(&bookCreateDTO)
// 	if errDTO != nil {
// 		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
// 		context.JSON(http.StatusBadRequest, res)
// 	} else {
// 		authHeader := context.GetHeader("Authorization")
// 		userID := c.getUserIDByToken(authHeader)
// 		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
// 		if err == nil {
// 			bookCreateDTO.UserID = convertedUserID
// 		}
// 		result := c.bookService.Insert(bookCreateDTO)
// 		response := helper.BuildResponse(true, "OK", result)
// 		context.JSON(http.StatusCreated, response)
// 	}
// }

// func (c *bookController) Update(context *gin.Context) {
// 	var bookUpdateDTO dto.BookUpdateDTO
// 	errDTO := context.ShouldBind(&bookUpdateDTO)
// 	if errDTO != nil {
// 		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
// 		context.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	authHeader := context.GetHeader("Authorization")
// 	token, errToken := c.jwtService.ValidateToken(authHeader)
// 	if errToken != nil {
// 		panic(errToken.Error())
// 	}
// 	claims := token.Claims.(jwt.MapClaims)
// 	userID := fmt.Sprintf("%v", claims["user_id"])
// 	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
// 		id, errID := strconv.ParseUint(userID, 10, 64)
// 		if errID == nil {
// 			bookUpdateDTO.UserID = id
// 		}
// 		result := c.bookService.Update(bookUpdateDTO)
// 		response := helper.BuildResponse(true, "OK", result)
// 		context.JSON(http.StatusOK, response)
// 	} else {
// 		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
// 		context.JSON(http.StatusForbidden, response)
// 	}
// }

// func (c *bookController) Delete(context *gin.Context) {
// 	var book entity.Book
// 	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
// 	if err != nil {
// 		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
// 		context.JSON(http.StatusBadRequest, response)
// 	}
// 	book.ID = id
// 	authHeader := context.GetHeader("Authorization")
// 	token, errToken := c.jwtService.ValidateToken(authHeader)
// 	if errToken != nil {
// 		panic(errToken.Error())
// 	}
// 	claims := token.Claims.(jwt.MapClaims)
// 	userID := fmt.Sprintf("%v", claims["user_id"])
// 	if c.bookService.IsAllowedToEdit(userID, book.ID) {
// 		c.bookService.Delete(book)
// 		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
// 		context.JSON(http.StatusOK, res)
// 	} else {
// 		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
// 		context.JSON(http.StatusForbidden, response)
// 	}
// }

// func (c *bookController) getUserIDByToken(token string) string {
// 	aToken, err := c.jwtService.ValidateToken(token)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	claims := aToken.Claims.(jwt.MapClaims)
// 	id := fmt.Sprintf("%v", claims["user_id"])
// 	return id
// }