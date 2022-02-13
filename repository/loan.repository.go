package repository

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/gorm"
)

//LoanRepository is contract what LoanRepository can do to db
type LoanRepository interface {
	InsertLoan(loan entity.Loan) entity.Loan
	AllLoans() []entity.Loan
	UpdateLoan(loan entity.Loan) entity.Loan
	FindLoanId(loanID string) entity.Loan
	DeleteLoan(loan entity.Loan)
}

type loanConnection struct {
	connection *gorm.DB
}

//NewLoanRepository is creates a new instance of LoanRepository
func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanConnection{
		connection: db,
	}
}

func (db *loanConnection) AllLoans() []entity.Loan {
	var loans []entity.Loan
	db.connection.Find(&loans)
	return loans
}

func (db *loanConnection) InsertLoan(loan entity.Loan) entity.Loan {
	db.connection.Create(&loan)
	db.connection.Find(&loan)
	return loan
}

func (db *loanConnection) FindLoanId(loanID string) entity.Loan {
	var loan entity.Loan
	db.connection.First(&loan, "id_loan = ?", loanID)
	return loan
}

func (db *loanConnection) UpdateLoan(loan entity.Loan) entity.Loan {
	db.connection.Updates(&loan)
	db.connection.Find(&loan)
	return loan
}

func (db *loanConnection) DeleteLoan(loan entity.Loan) {
	db.connection.Delete(&loan)
}