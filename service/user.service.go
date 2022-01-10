package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/repository"
)

//UserService is a contract.....
type UserService interface {
	All() []entity.User
	Delete(user entity.User)
	Insert(u dto.UserCreateDTO) entity.User
	// Update(user dto.UserUpdateDTO) entity.User
	// Profile(userID string) entity.User
	FindByID(userID int64) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Insert(u dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{}
	fmt.Println("userToCreate inisiasi", userToCreate)
	fmt.Println("==========")

	// Mengisi variable userToCreate
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&u))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	// fmt.Println("userToCreate sblm insertuser", userToCreate)
	// fmt.Println("==========")
	res := service.userRepository.InsertUser(userToCreate)
	// fmt.Println("res setelah insertuser", res)
	// fmt.Println("==========")
	// os.Exit(1)
	// res := service.userRepository.InsertUser(userToCreate)
	
	return res
}



func (service *userService) All() []entity.User {
	return service.userRepository.AllUser()
}

func (service *userService) Delete(user entity.User) {
	service.userRepository.DeleteUser(user)
}

func (service *userService) FindByID(userID int64) entity.User {
	return service.userRepository.FindUserID(userID)
}


// func (service *bookService) Delete(b entity.Book) {
// 	service.bookRepository.DeleteBook(b)
// }




// func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
// 	userToUpdate := entity.User{}
// 	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
// 	if err != nil {
// 		log.Fatalf("Failed map %v:", err)
// 	}
// 	updatedUser := service.userRepository.UpdateUser(userToUpdate)
// 	return updatedUser
// }

// func (service *userService) Profile(userID string) entity.User {
// 	return service.userRepository.ProfileUser(userID)
// }
