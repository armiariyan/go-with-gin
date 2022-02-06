package repository

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type BorrowerRepository interface {
	
	// FindByEmail(email string) entity.User
	// ProfileUser(userID string) entity.User
	InsertBorrower(borrower entity.Borrower) entity.Borrower
	AllBorrowers() []entity.Borrower
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

func (db *borrowerConnection) AllBorrowers() []entity.Borrower {
	var borrowers []entity.Borrower
	db.connection.Preload("User").Preload("Opt_house").Find(&borrowers)
	return borrowers
}

func (db *borrowerConnection) InsertBorrower(borrower entity.Borrower) entity.Borrower {
	db.connection.Create(&borrower)
	db.connection.Preload("User").Preload("Opt_house").Find(&borrower)
	return borrower
}
