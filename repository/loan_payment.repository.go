package repository

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/gorm"
)

//Loan_paymentRepository is contract what Loan_paymentRepository can do to db
type Loan_paymentRepository interface {
	InsertLoan_payment(loan_payment entity.Loan_payment) entity.Loan_payment
	AllLoan_payments() []entity.Loan_payment
	UpdateLoan_payment(loan_payment entity.Loan_payment) entity.Loan_payment
	FindLoan_paymentId(loan_paymentID string) entity.Loan_payment
	DeleteLoan_payment(loan_payment entity.Loan_payment)
}

type loan_paymentConnection struct {
	connection *gorm.DB
}

//NewLoan_paymentRepository is creates a new instance of Loan_paymentRepository
func NewLoan_paymentRepository(db *gorm.DB) Loan_paymentRepository {
	return &loan_paymentConnection{
		connection: db,
	}
}

func (db *loan_paymentConnection) AllLoan_payments() []entity.Loan_payment {
	var loan_payments []entity.Loan_payment
	db.connection.Find(&loan_payments)
	return loan_payments
}

func (db *loan_paymentConnection) InsertLoan_payment(loan_payment entity.Loan_payment) entity.Loan_payment {
	db.connection.Create(&loan_payment)
	db.connection.Find(&loan_payment)
	return loan_payment
}

func (db *loan_paymentConnection) FindLoan_paymentId(loan_paymentID string) entity.Loan_payment {
	var loan_payment entity.Loan_payment
	db.connection.First(&loan_payment, "id_payment = ?", loan_paymentID)
	return loan_payment
}

func (db *loan_paymentConnection) UpdateLoan_payment(loan_payment entity.Loan_payment) entity.Loan_payment {
	db.connection.Updates(&loan_payment)
	db.connection.Find(&loan_payment)
	return loan_payment
}

func (db *loan_paymentConnection) DeleteLoan_payment(loan_payment entity.Loan_payment) {
	db.connection.Delete(&loan_payment)
}