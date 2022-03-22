package repository

import (
	"log"

	"gitlab.com/armiariyan/intern_golang/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	
	// FindByEmail(email string) entity.User
	// ProfileUser(userID string) entity.User
	InsertUser(user entity.User) entity.User
	AllUser() []entity.User
	DeleteUser(user entity.User)
	UpdateUser(user entity.User) entity.User
	FindUserID(userID int64) entity.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	VerifyCredential(email string, password string) interface{}
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Create(&user)
	// db.connection.Preload("User").Find(&b)
	return user
}

func (db *userConnection) AllUser() []entity.User {
	var users []entity.User
	db.connection.Find(&users)
	return users
}

func (db *userConnection) DeleteUser(user entity.User) {
	db.connection.Delete(&user)
}

func (db *userConnection) FindUserID(userID int64) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}


func (db *userConnection) UpdateUser(user entity.User) entity.User {	
	// For hash new password
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {		
		tempUser := db.FindUserID(user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Updates(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func hashAndSalt(pwd []byte) string {	
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}




// func (db *userConnection) FindByEmail(email string) entity.User {
// 	var user entity.User
// 	db.connection.Where("email = ?", email).Take(&user)
// 	return user
// }

// func (db *userConnection) ProfileUser(userID string) entity.User {
// 	var user entity.User
// 	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)
// 	return user
// }
