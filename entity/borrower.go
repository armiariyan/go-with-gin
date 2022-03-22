package entity

type Borrower struct {
	ID          int64     `gorm:"primary_key:auto_increment" json:"id"`
	Id_borrower string    `gorm:"type:varchar(255)" json:"id_borrower"`
	Id_user     int64     `json:"id_user"`
	House       int64     `json:"house"`
	User        User      `gorm:"foreignKey:Id_user" json:"user"`    //Belongs to one user
	Opt_house   Opt_house `gorm:"foreignKey:House" json:"opt_house"` //Belongs to one house
}
