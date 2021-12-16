package entities

type Account struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Type 	 int 	
}