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
	// AllUser() []entity.User
	// DeleteUser(user entity.User)
	// UpdateUser(user entity.User) entity.User
	// FindUserID(userID int64) entity.User
	// IsDuplicateEmail(email string) (tx *gorm.DB)
	// VerifyCredential(email string, password string) interface{}
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

func (db *lenderConnection) InsertLender(lender entity.Lender) entity.Lender {
	// fmt.Println("lender", lender)
	db.connection.Create(&lender)
	// fmt.Println("lender after create", lender)
	db.connection.Preload("User").Preload("House").Find(&lender)
	// fmt.Println("lender after preload", lender)

	return lender
}
