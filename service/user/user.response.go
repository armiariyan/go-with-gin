package user

import (
	"intern_golang/entities"
)

type UserResponse struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Id_number  string `json:"id_number"`
	Type       int    `json:"type"`
	// Token string `json:"token,omitempty"`
}

func NewUserResponse(user entities.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Username:  user.Username,
		Password: user.Password,
		Email: user.Email,
		First_name: user.Email,
		Last_name: user.Last_name,
		Id_number: user.Id_number,
		Type: user.Type,
	}
}
