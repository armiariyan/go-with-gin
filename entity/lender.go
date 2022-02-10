package entity

type Lender struct {
	Id_lender   string `gorm:"primary_key;type:varchar(255)" json:"id_lender"`
	Id_user     int64  `json:"id_user"`
	Sumber_dana string `gorm:"type:varchar(255)" json:"sumber_dana"`
	User        User   `gorm:"foreignKey:Id_user" json:"user"` //Belongs to one user
}
