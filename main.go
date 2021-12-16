package main

import (
	"intern_golang/config"
	"intern_golang/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db 	*gorm.DB = config.Setup_db_connection()
	userRepo repo.UserRepository = repo.NewProductRepo(db)
	userService 

)

func main() {
	defer config.Close_db_connection(db)
	server := gin.Default()

	// user_routes := server.Group("api/users/")
	// {
	// 	user_routes.GET("/", usersHandler.GetAll)
	// 	// user_routes.POST("/", usersHandler.CreateUser)
	// }
	
	server.GET("/", rootHandler)
	
	server.Run()

}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Armia Riyan",
	})
}



