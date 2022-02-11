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

//RequestLoanService is a contract about something that this service can do
type RequestLoanService interface {
	CreateRequestLoan(requestLoan dto.RequestLoanCreateDTO) entity.Request_loan
	CreateIdRequestLoan() string
	All() []entity.Request_loan
	FindByID(requestLoanID string) entity.Request_loan
	Update(b dto.RequestLoanUpdateDTO) entity.Request_loan
	Delete(requestLoans entity.Request_loan)
}

type requestLoanService struct {
	requestLoanRepository repository.RequestLoanRepository
}

//NewRequestLoanService creates a new instance of AuthService
func NewRequestLoanService(requestLoanRep repository.RequestLoanRepository) RequestLoanService {
	return &requestLoanService{
		requestLoanRepository: requestLoanRep,
	}
}


func (service *requestLoanService) CreateIdRequestLoan() string {
	rand.Seed(time.Now().UnixNano())
	id_requestLoan := "LOAN-" + RandomLetters(6) + "-" + GetCurrentTime()
	return id_requestLoan
}

func (service *requestLoanService) CreateRequestLoan(requestLoan dto.RequestLoanCreateDTO) entity.Request_loan {
	requestLoanToCreate := entity.Request_loan{}
	
	err := smapping.FillStruct(&requestLoanToCreate, smapping.MapFields(&requestLoan))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	
	res := service.requestLoanRepository.InsertRequestLoan(requestLoanToCreate)
	return res
}

func (service *requestLoanService) FindByID(requestLoanID string) entity.Request_loan {
	return service.requestLoanRepository.FindRequestLoanId(requestLoanID)
}

func (service *requestLoanService) All() []entity.Request_loan {
	return service.requestLoanRepository.AllRequestLoans()
}

func (service *requestLoanService) Update(requestLoan dto.RequestLoanUpdateDTO) entity.Request_loan {
	requestLoanToUpdate := entity.Request_loan{}
	// Fill the variable
	err := smapping.FillStruct(&requestLoanToUpdate, smapping.MapFields(&requestLoan))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	// Update the variable
	updatedRequestLoan := service.requestLoanRepository.UpdateRequestLoan(requestLoanToUpdate)
	return updatedRequestLoan
}

func (service *requestLoanService) Delete(requestLoan entity.Request_loan) {
	service.requestLoanRepository.DeleteRequestLoan(requestLoan)
}



