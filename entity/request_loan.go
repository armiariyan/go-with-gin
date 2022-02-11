package entity

type Request_loan struct {
	Id_loan               string                `gorm:"primary_key;type:varchar(255)" json:"id_loan"`
	Id_borrower           string                `gorm:"type:varchar(255)" json:"id_borrower"`
	Loan_name             string                `gorm:"type:varchar(255)" json:"loan_name"`
	Loan_amount           int64                 `json:"loan_amount"`
	Loan_duration         int64                 `json:"loan_duration"`
	Payment_frequency     int64                 `json:"payment_frequency"`
	Payment_type          int64                 `json:"payment_type"`
	Status                string                `gorm:"type:varchar(124)" json:"status"`
	Opt_payment_frequency Opt_payment_frequency `gorm:"foreignKey:Payment_frequency" json:"opt_payment_frequency"`
	Opt_payment_type      Opt_payment_type      `gorm:"foreignKey:Payment_type" json:"opt_payment_type"`
}
