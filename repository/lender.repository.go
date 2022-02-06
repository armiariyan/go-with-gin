package repository

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type LenderRepository interface {
	
	// FindByEmail(email string) entity.User
	// ProfileUser(userID string) entity.User
	InsertLender(lender entity.Lender) entity.Lender
	AllLenders() []entity.Lender
	UpdateLender(lender entity.Lender) entity.Lender
	FindLenderId(lenderID string) entity.Lender
	DeleteLender(lender entity.Lender)
}

type lenderConnection struct {
	connection *gorm.DB
}

//NewLenderRepository is creates a new instance of LenderRepository
func NewLenderRepository(db *gorm.DB) LenderRepository {
	return &lenderConnection{
		connection: db,
	}
}

func (db *lenderConnection) AllLenders() []entity.Lender {
	var lenders []entity.Lender
	db.connection.Preload("User").Find(&lenders)
	return lenders
}

func (db *lenderConnection) InsertLender(lender entity.Lender) entity.Lender {
	db.connection.Create(&lender)
	db.connection.Preload("User").Find(&lender)
	return lender
}

func (db *lenderConnection) FindLenderId(lenderID string) entity.Lender {
	var lender entity.Lender
	db.connection.First(&lender, "id_lender = ?", lenderID)
	return lender
}

func (db *lenderConnection) UpdateLender(lender entity.Lender) entity.Lender {
	db.connection.Updates(&lender)
	db.connection.Preload("User").Find(&lender)
	return lender
}

func (db *lenderConnection) DeleteLender(lender entity.Lender) {
	db.connection.Delete(&lender)
}