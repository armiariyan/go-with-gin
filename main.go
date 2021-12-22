package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/armiariyan/intern_golang/config"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello armia",
		})
	})
	r.Run()
}