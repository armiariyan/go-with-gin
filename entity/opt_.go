package entity

type Opt_house struct {
	ID    int64  `gorm:"primary_key:auto_increment" json:"id"`
	Label string `gorm:"type:varchar(255)" json:"label"`
}

type Opt_payment_frequency struct {
	ID    int64  `gorm:"primary_key:auto_increment" json:"id"`
	Label string `gorm:"type:varchar(255)" json:"label"`
}

type Opt_payment_type struct {
	ID    int64  `gorm:"primary_key:auto_increment" json:"id"`
	Label string `gorm:"type:varchar(255)" json:"label"`
}

type Opt_status struct {
	ID    int64  `gorm:"primary_key:auto_increment" json:"id"`
	Label string `gorm:"type:varchar(255)" json:"label"`
}
