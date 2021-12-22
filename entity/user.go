package entity

type User struct {
	ID         int64  `gorm:"primary_key:auto_increment" json:"id"`
	Username   string `gorm:"type:varchar(255)" json:"username"`
	Password   string `gorm:"type:varchar(255)" json:"password"`
	Email      string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	First_name string `gorm:"type:varchar(255);" json:"first_name"`
	Last_name  string `gorm:"type:varchar(255)" json:"last_name"`
	Id_number  string `gorm:"type:varchar(255)" json:"id_number"`
	Type       int    `json:"type"`
}