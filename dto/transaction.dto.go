package dto

import (
	"time"
)

//TransactionUpdateDTO is used by client when PUT update transaction data
type TransactionUpdateDTO struct {
	Id_transaction        string                `json:"id_transaction" form:"id_transaction"`
	Id_borrower           string                `json:"id_borrower" form:"id_borrower"`
	Id_lender             string                `json:"id_lender" form:"id_lender"`
	Lend_number           int64                 `json:"lend_number" form:"lend_number"`
	Interest_rate         float64               `json:"interest_rate" form:"interest_rate"`
	Borrower_deal_date    time.Time              `json:"borrower_deal_date" form:"borrower_deal_date"`
	Due_date          	  time.Time              `json:"due_date" form:"due_date"`
}

//TransactionCreateDTO is is a model that client use to create transaction data
type TransactionCreateDTO struct {
	Id_transaction        string                `json:"id_transaction" form:"id_transaction"`
	Id_borrower           string                `json:"id_borrower" form:"id_borrower"`
	Id_lender             string                `json:"id_lender" form:"id_lender"`
	Lend_number           int64                 `json:"lend_number" form:"lend_number"`
	Interest_rate         float64               `json:"interest_rate" form:"interest_rate"`
	Borrower_deal_date    time.Time             `json:"borrower_deal_date" form:"borrower_deal_date"`
	Due_date          	  time.Time             `json:"due_date" form:"due_date"`
}
