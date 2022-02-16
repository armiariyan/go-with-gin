package entity

import "time"

type Loan struct {
	Id_loan     	 string    	 `gorm:"primary_key;type:varchar(255)" json:"id_loan"`
	Id_lender   	 string    	 `gorm:"type:varchar(255)" json:"id_lender"`
	Id_borrower 	 string		 `gorm:"type:varchar(255)" json:"id_borrower"`
	Deal_date   	 time.Time 	 `json:"deal_date"`
	Va_lender  		 string    	 `json:"va_lender"`
	Loan_amount		 int64		 `json:"loan_amount"`
	Confirmed_amount int64 		 `json:"confirmed_amount"`
	Status			 int64		 `json:"status"`
	Opt_status 		 Opt_status	 `gorm:"foreignKey:Status" json:"opt_status"`
}
