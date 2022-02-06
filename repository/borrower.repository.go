package repository

import (
	"fmt"

	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type BorrowerRepository interface {
	
	// FindByEmail(email string) entity.User
	// ProfileUser(userID string) entity.User
	InsertBorrower(borrower entity.Borrower) entity.Borrower
	// AllUser() []entity.User
	// DeleteUser(user entity.User)
	// UpdateUser(user entity.User) entity.User
	// FindUserID(userID int64) entity.User
	// IsDuplicateEmail(email string) (tx *gorm.DB)
	// VerifyCredential(email string, password string) interface{}
}

type borrowerConnection struct {
	connection *gorm.DB
}

//NewBorrowerRepository is creates a new instance of BorrowerRepository
func NewBorrowerRepository(db *gorm.DB) BorrowerRepository {
	return &borrowerConnection{
		connection: db,
	}
}

func (db *borrowerConnection) InsertBorrower(borrower entity.Borrower) entity.Borrower {
	fmt.Println("borrower", borrower)
	db.connection.Create(&borrower)
	fmt.Println("borrower after create", borrower)
	db.connection.Preload("User").Preload("House").Find(&borrower)
	fmt.Println("borrower after preload", borrower)

	return borrower
}
