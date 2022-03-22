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

//Loan_paymentService is a contract about something that this service can do
type Loan_paymentService interface {
	CreateLoan_payment(loan_payment dto.Loan_paymentCreateDTO) entity.Loan_payment
	CreateIdLoan_payment() string
	All() []entity.Loan_payment
	FindByID(loan_paymentID string) entity.Loan_payment
	Update(b dto.Loan_paymentUpdateDTO) entity.Loan_payment
	Delete(loan_payments entity.Loan_payment)
}

type loan_paymentService struct {
	loan_paymentRepository repository.Loan_paymentRepository
}

//NewLoan_paymentService creates a new instance of AuthService
func NewLoan_paymentService(loan_paymentRep repository.Loan_paymentRepository) Loan_paymentService {
	return &loan_paymentService{
		loan_paymentRepository: loan_paymentRep,
	}
}


func (service *loan_paymentService) CreateIdLoan_payment() string {
	rand.Seed(time.Now().UnixNano())
	id_loan_payment := "PAY-" + GetCurrentTime() + "-" + RandomLetters(4)
	return id_loan_payment
}

func (service *loan_paymentService) CreateLoan_payment(loan_payment dto.Loan_paymentCreateDTO) entity.Loan_payment {
	loan_paymentToCreate := entity.Loan_payment{}
	
	err := smapping.FillStruct(&loan_paymentToCreate, smapping.MapFields(&loan_payment))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	
	res := service.loan_paymentRepository.InsertLoan_payment(loan_paymentToCreate)
	return res
}

func (service *loan_paymentService) FindByID(loan_paymentID string) entity.Loan_payment {
	return service.loan_paymentRepository.FindLoan_paymentId(loan_paymentID)
}

func (service *loan_paymentService) All() []entity.Loan_payment {
	return service.loan_paymentRepository.AllLoan_payments()
}

func (service *loan_paymentService) Update(loan_payment dto.Loan_paymentUpdateDTO) entity.Loan_payment {
	loan_paymentToUpdate := entity.Loan_payment{}
	// Fill the variable
	err := smapping.FillStruct(&loan_paymentToUpdate, smapping.MapFields(&loan_payment))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	// Update the variable
	updatedLoan_payment := service.loan_paymentRepository.UpdateLoan_payment(loan_paymentToUpdate)
	return updatedLoan_payment
}

func (service *loan_paymentService) Delete(loan_payment entity.Loan_payment) {
	service.loan_paymentRepository.DeleteLoan_payment(loan_payment)
}



