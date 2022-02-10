package dto

//RequestLoanUpdateDTO is used by client when PUT update lender data
type RequestLoanUpdateDTO struct {
	Id_loan               string `json:"id_loan" form:"id_loan"`
	Id_borrower           int64  `json:"id_borrower" form:"id_borrower"`
	Loan_name             string `json:"loan_name" form:"loan_name" binding:"required"`
	Loan_amount           int64  `json:"loan_amount" form:"loan_amount" binding:"required"`
	Loan_duration         int64  `json:"loan_duration" form:"loan_duration" binding:"required"`
	Payment_frequency     int8   `json:"payment_frequency" form:"payment_frequency" binding:"required"`
	Payment_type          int8   `json:"payment_type" form:"payment_type" binding:"required"`
	Status				  string `json:"status" form:"status" binding:"oneof=waiting admin completed declined"`
}

//RequestLoanCreateDTO is is a model that client use to create lender data
type RequestLoanCreateDTO struct {
	Id_loan               string `json:"id_loan" form:"id_loan"`
	Id_borrower           int64  `json:"id_borrower" form:"id_borrower"`
	Loan_name             string `json:"loan_name" form:"loan_name" binding:"required"`
	Loan_amount           int64  `json:"loan_amount" form:"loan_amount" binding:"required"`
	Loan_duration         int64  `json:"loan_duration" form:"loan_duration" binding:"required"`
	Payment_frequency     int8   `json:"payment_frequency" form:"payment_frequency" binding:"required"`
	Payment_type          int8   `json:"payment_type" form:"payment_type" binding:"required"`
	Status				  string `json:"status" form:"status" binding:"oneof=waiting admin completed declined"`
}