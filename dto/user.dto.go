package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID         int64  `json:"id" form:"id"`
	Username   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	First_name string `json:"first_name" form:"first_name" binding:"required"`
	Last_name  string `json:"last_name" form:"last_name"`
	Id_number  string `json:"id_number" form:"id_number" binding:"required"`
	Type       int    `json:"type" form:"type" binding:"required,oneof=1 2"`
}

//UserCreateDTO is is a model that clinet use when create a new user
type UserCreateDTO struct {
	Username   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	First_name string `json:"first_name" form:"first_name" binding:"required"`
	Last_name  string `json:"last_name" form:"last_name"`
	Id_number  string `json:"id_number" form:"id_number" binding:"required"`
	Type       int    `json:"type" form:"type" binding:"required,oneof=1 2"`
}
