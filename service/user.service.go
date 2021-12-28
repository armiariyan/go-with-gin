package service

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/repository"
)

//UserService is a contract.....
type UserService interface {
	All() []entity.User
	// Update(user dto.UserUpdateDTO) entity.User
	// Profile(userID string) entity.User
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
