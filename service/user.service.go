package service

import (
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
	Update(u dto.UserUpdateDTO) entity.User
	// Profile(userID string) entity.User
	FindByID(userID int64) entity.User
	// IsDuplicateEmail(email string) bool
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


func (service *userService) All() []entity.User {
	return service.userRepository.AllUser()
}

func (service *userService) Delete(user entity.User) {
	service.userRepository.DeleteUser(user)
}

func (service *userService) FindByID(userID int64) entity.User {
	return service.userRepository.FindUserID(userID)
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	// Fill the variable
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	// Update the variable
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

// func (service *userService) IsDuplicateEmail(email string) bool {
// 	fmt.Println("email in user service", email)
// 	fmt.Println("=================")
// 	res := service.userRepository.IsDuplicateEmail(email)
// 	return !(res.Error == nil)
// }


// func (service *userService) Profile(userID string) entity.User {
// 	return service.userRepository.ProfileUser(userID)
// }
