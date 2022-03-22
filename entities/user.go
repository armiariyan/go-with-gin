package entities

type User struct {
	ID         int64  `gorm:"primary_key:auto_increment" json:"-"`
	Username   string `gorm:"type:varchar(255)" json:"-"`
	Password   string `gorm:"type:varchar(255)" json:"-"`
	Email      string `gorm:"uniqueIndex;type:varchar(255)" json:"-"`
	First_name string `gorm:"type:varchar(255);" json:"-"`
	Last_name  string `gorm:"type:varchar(255)" json:"-"`
	Id_number  string `gorm:"type:varchar(255)" json:"-"`
	Type       int    `json:"-"`
}