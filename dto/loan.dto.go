package dto

import (
	"time"
)

//LoanUpdateDTO is used by client when PUT update loan data
type LoanUpdateDTO struct {
	Id_loan     	 string     `gorm:"type:varchar(255)" json:"id_loan" binding:"required" form:"id_loan"`
	Id_lender   	 string     `gorm:"type:varchar(255)" json:"id_lender" binding:"required" form:"id_lender"`
	Id_borrower 	 string     `gorm:"type:varchar(255)" json:"id_borrower" binding:"required" form:"id_borrower"`
	Deal_date   	 time.Time  `json:"deal_date" binding:"required" form:"deal_date"`
	Va_lender  		 string     `json:"va_lender" binding:"required,max=16" form:"va_lender"`
	Confirmed_amount int64 		`json:"confirmed_amont" binding:"required" form:"confirmed_amont"`
	Status			 int64		`json:"status" binding:"required" form:"status"`
}

//LoanCreateDTO is is a model that client use to create loan data
type LoanCreateDTO struct {
	Id_loan     	 string     `gorm:"type:varchar(255)" json:"id_loan" binding:"required" form:"id_loan"`
	Id_lender   	 string     `gorm:"type:varchar(255)" json:"id_lender" binding:"required" form:"id_lender"`
	Id_borrower 	 string     `gorm:"type:varchar(255)" json:"id_borrower" binding:"required" form:"id_borrower"`
	Deal_date   	 time.Time  `json:"deal_date" binding:"required" form:"deal_date"`
	Va_lender  		 string     `json:"va_lender" binding:"required,max=16" form:"va_lender"`
	Confirmed_amount int64 		`json:"confirmed_amont" binding:"required" form:"confirmed_amont"`
	Status			 int64		`json:"status" binding:"required" form:"status"`
}
