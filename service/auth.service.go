package service

import (
	"log"

	"github.com/mashingan/smapping"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/repository"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	
	// fmt.Println("password =", password)
	// fmt.Println("=====================")
	if v, ok := res.(entity.User); ok {
		// fmt.Println("v =", v.Password)
		// fmt.Println("=====================")
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	// fmt.Println("hashedPwd =", hashedPwd)
	// fmt.Println("=====================")
	byteHash := []byte(hashedPwd)
	// fmt.Println("plain password =", plainPassword)
	// fmt.Println("=====================")
	// fmt.Println("byteHash =", byteHash)
	// fmt.Println("=====================")
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
