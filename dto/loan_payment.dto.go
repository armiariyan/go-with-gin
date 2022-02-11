package dto

import (
	"time"
)

//Loan_paymentUpdateDTO is used by client when PUT update transaction data
type Loan_paymentUpdateDTO struct {
	Id_payment      string      `form:"id_payment" json:"id_payment"`
	Id_transaction  string      `form:"id_transaction" json:"id_transaction"`
	Id_lender       string      `form:"id_lender" json:"id_lender"`
	Id_borrower     string      `form:"id_borrower" json:"id_borrower"`
	Id_loan         string      `form:"id_loan" json:"id_loan"`
	Payment_date	time.Time	`json:"payment_date"`
}

//Loan_paymentCreateDTO is is a model that client use to create transaction data
type Loan_paymentCreateDTO struct {
	Id_payment      string      `form:"id_payment" json:"id_payment"`
	Id_transaction  string      `form:"id_transaction" json:"id_transaction"`
	Id_lender       string      `form:"id_lender" json:"id_lender"`
	Id_borrower     string      `form:"id_borrower" json:"id_borrower"`
	Id_loan         string      `form:"id_loan" json:"id_loan"`
	Payment_date	time.Time	`json:"payment_date"`
}
