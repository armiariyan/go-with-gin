package service

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/repository"
)

//UserService is a contract.....
type UserService interface {
	All() []entity.User
	Delete(b entity.User)
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

func (service *userService) All() []entity.User {
	return service.userRepository.AllUser()
}

func (service *userService) Delete(b entity.User) {
	service.userRepository.DeleteUser(b)
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
