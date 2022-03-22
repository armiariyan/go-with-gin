package repository

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type RequestLoanRepository interface {
	InsertRequestLoan(requestLoan entity.Request_loan) entity.Request_loan
	AllRequestLoans() []entity.Request_loan
	UpdateRequestLoan(requestLoan entity.Request_loan) entity.Request_loan
	FindRequestLoanId(requestLoanID string) entity.Request_loan
	DeleteRequestLoan(requestLoan entity.Request_loan)
}

type requestLoanConnection struct {
	connection *gorm.DB
}

//NewRequestLoanRepository is creates a new instance of RequestLoanRepository
func NewRequestLoanRepository(db *gorm.DB) RequestLoanRepository {
	return &requestLoanConnection{
		connection: db,
	}
}

func (db *requestLoanConnection) AllRequestLoans() []entity.Request_loan {
	var requestLoans []entity.Request_loan
	db.connection.Preload("Opt_payment_frequency").Preload("Opt_payment_type").Find(&requestLoans)
	return requestLoans
}

func (db *requestLoanConnection) InsertRequestLoan(requestLoan entity.Request_loan) entity.Request_loan {
	db.connection.Create(&requestLoan)
	db.connection.Preload("Opt_payment_frequency").Preload("Opt_payment_type").Find(&requestLoan)
	return requestLoan
}

func (db *requestLoanConnection) FindRequestLoanId(requestLoanID string) entity.Request_loan {
	var requestLoan entity.Request_loan
	db.connection.First(&requestLoan, "id_loan = ?", requestLoanID)
	return requestLoan
}

func (db *requestLoanConnection) UpdateRequestLoan(requestLoan entity.Request_loan) entity.Request_loan {
	db.connection.Updates(&requestLoan)
	db.connection.Preload("Opt_payment_frequency").Preload("Opt_payment_type").Find(&requestLoan)
	return requestLoan
}

func (db *requestLoanConnection) DeleteRequestLoan(requestLoan entity.Request_loan) {
	db.connection.Delete(&requestLoan)
}