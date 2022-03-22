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
	UpdateBorrower(borrower entity.Borrower) entity.Borrower
	FindBorrowerId(borrowerID string) entity.Borrower
	DeleteBorrower(borrower entity.Borrower)
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

func (db *borrowerConnection) FindBorrowerId(borrowerID string) entity.Borrower {
	var borrower entity.Borrower
	db.connection.First(&borrower, "id_borrower = ?", borrowerID)
	return borrower
}

func (db *borrowerConnection) UpdateBorrower(borrower entity.Borrower) entity.Borrower {
	db.connection.Updates(&borrower)
	db.connection.Preload("User").Preload("Opt_house").Find(&borrower)
	return borrower
}

func (db *borrowerConnection) DeleteBorrower(borrower entity.Borrower) {
	db.connection.Delete(&borrower)
}