package service

import (
	"log"
	"math/rand"
	"time"

	"github.com/mashingan/smapping"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/repository"
)

//LoanService is a contract about something that this service can do
type LoanService interface {
	CreateLoan(loan dto.LoanCreateDTO) entity.Loan
	CreateIdLoan() string
	All() []entity.Loan
	FindByID(loanID string) entity.Loan
	Update(b dto.LoanUpdateDTO) entity.Loan
	Delete(loans entity.Loan)
}

type loanService struct {
	loanRepository repository.LoanRepository
}

//NewLoanService creates a new instance of AuthService
func NewLoanService(loanRep repository.LoanRepository) LoanService {
	return &loanService{
		loanRepository: loanRep,
	}
}


func (service *loanService) CreateIdLoan() string {
	rand.Seed(time.Now().UnixNano())
	id_loan := "PAY-" + GetCurrentTime() + "-" + RandomLetters(4)
	return id_loan
}

func (service *loanService) CreateLoan(loan dto.LoanCreateDTO) entity.Loan {
	loanToCreate := entity.Loan{}
	
	err := smapping.FillStruct(&loanToCreate, smapping.MapFields(&loan))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	
	res := service.loanRepository.InsertLoan(loanToCreate)
	return res
}

func (service *loanService) FindByID(loanID string) entity.Loan {
	return service.loanRepository.FindLoanId(loanID)
}

func (service *loanService) All() []entity.Loan {
	return service.loanRepository.AllLoans()
}

func (service *loanService) Update(loan dto.LoanUpdateDTO) entity.Loan {
	loanToUpdate := entity.Loan{}
	// Fill the variable
	err := smapping.FillStruct(&loanToUpdate, smapping.MapFields(&loan))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	// Update the variable
	updatedLoan := service.loanRepository.UpdateLoan(loanToUpdate)
	return updatedLoan
}

func (service *loanService) Delete(loan entity.Loan) {
	service.loanRepository.DeleteLoan(loan)
}



