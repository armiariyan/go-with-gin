package entity

import (
	"time"
)


type Loan_payment struct {
	ID              int64 	    `gorm:"primary_key:auto_increment" json:"id"`
	Id_payment      string      `gorm:"type:varchar(255)" json:"id_payment"`
	Id_transaction  string      `gorm:"type:varchar(255)" json:"id_transaction"`
	Id_lender       string      `gorm:"type:varchar(255)" json:"id_lender"`
	Id_borrower     string      `gorm:"type:varchar(255)" json:"id_borrower"`
	Id_loan         string      `gorm:"type:varchar(255)" json:"id_loan"`
	Payment_date	time.Time	`json:"payment_date"`
}

