package entity

type Lender struct {
	ID          int64  `gorm:"primary_key:auto_increment" json:"id"`
	Id_lender   string `gorm:"type:varchar(255)" json:"id_lender"`
	Id_user     int64  `json:"id_user"`
	Sumber_dana string `gorm:"type:varchar(255)" json:"sumber_dana"`
	User        User   `gorm:"foreignKey:Id_user" json:"user"` //Belongs to one user
}
