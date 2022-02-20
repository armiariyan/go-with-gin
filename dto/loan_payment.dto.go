package dto

import (
	"time"
)

//Loan_paymentUpdateDTO is used by client when PUT update transaction data
type Loan_paymentUpdateDTO struct {
	ID         		int64		`json:"id" form:"id"`
	Id_payment      string      `form:"id_payment" json:"id_payment" binding:"required"`
	Id_transaction  string      `form:"id_transaction" json:"id_transaction" binding:"required"`
	Id_lender       string      `form:"id_lender" json:"id_lender" binding:"required"`
	Id_borrower     string      `form:"id_borrower" json:"id_borrower" binding:"required"`
	Id_loan         string      `form:"id_loan" json:"id_loan" binding:"required"`
	Payment_date	time.Time	`json:"payment_date" binding:"required"`
}

//Loan_paymentCreateDTO is is a model that client use to create transaction data
type Loan_paymentCreateDTO struct {
	Id_payment      string      `form:"id_payment" json:"id_payment" binding:"required"`
	Id_transaction  string      `form:"id_transaction" json:"id_transaction" binding:"required"`
	Id_lender       string      `form:"id_lender" json:"id_lender" binding:"required"`
	Id_borrower     string      `form:"id_borrower" json:"id_borrower" binding:"required"`
	Id_loan         string      `form:"id_loan" json:"id_loan" binding:"required"`
	Payment_date	time.Time	`json:"payment_date" binding:"required"`
}
