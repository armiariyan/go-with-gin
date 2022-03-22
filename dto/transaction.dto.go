package dto

import (
	"time"
)

//TransactionUpdateDTO is used by client when PUT update transaction data
type TransactionUpdateDTO struct {
	ID         			  int64			    	`json:"id" form:"id"`
	Id_transaction        string                `json:"id_transaction" form:"id_transaction" binding:"required"`
	Id_borrower           string                `json:"id_borrower" form:"id_borrower" binding:"required"`
	Id_lender             string                `json:"id_lender" form:"id_lender" binding:"required"`
	Lend_number           int64                 `json:"lend_number" form:"lend_number" binding:"required"`
	Interest_rate         float64               `json:"interest_rate" form:"interest_rate" binding:"required"`
	Borrower_deal_date    time.Time              `json:"borrower_deal_date" form:"borrower_deal_date" binding:"required"`
	Due_date          	  time.Time              `json:"due_date" form:"due_date" binding:"required"`
}

//TransactionCreateDTO is is a model that client use to create transaction data
type TransactionCreateDTO struct {
	Id_transaction        string                `json:"id_transaction" form:"id_transaction" binding:"required"`
	Id_borrower           string                `json:"id_borrower" form:"id_borrower" binding:"required"`
	Id_lender             string                `json:"id_lender" form:"id_lender" binding:"required"`
	Lend_number           int64                 `json:"lend_number" form:"lend_number" binding:"required"`
	Interest_rate         float64               `json:"interest_rate" form:"interest_rate" binding:"required"`
	Borrower_deal_date    time.Time             `json:"borrower_deal_date" form:"borrower_deal_date" binding:"required"`
	Due_date          	  time.Time             `json:"due_date" form:"due_date" binding:"required"`
}
