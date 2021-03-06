package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Opt_house{},&entity.Opt_payment_frequency{},&entity.Opt_payment_type{},&entity.Opt_status{},&entity.Lender{},&entity.Request_loan{}, &entity.Borrower{}, &entity.Transaction{}, &entity.Loan_payment{}, &entity.Loan{})

	// db.AutoMigrate(&entity.Request_loan{},&entity.Borrower{})
	return db
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
