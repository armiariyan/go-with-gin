package entity

import (
	"time"
)


type Transaction struct {
	ID          		  int64 			    `gorm:"primary_key:auto_increment" json:"id"`
	Id_transaction        string                `gorm:"type:varchar(255)" json:"id_transaction"`
	Id_borrower           string                `gorm:"type:varchar(255)" json:"id_borrower"`
	Id_lender             string                `gorm:"type:varchar(255)" json:"id_lender"`
	Lend_number           int64                 `json:"lend_number"`
	Interest_rate         float64               `json:"interest_rate"`
	Borrower_deal_date    time.Time             `json:"borrower_deal_date"`
	Due_date          	  time.Time             `json:"due_date"`
}

