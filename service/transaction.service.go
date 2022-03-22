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

//TransactionService is a contract about something that this service can do
type TransactionService interface {
	CreateTransaction(transaction dto.TransactionCreateDTO) entity.Transaction
	CreateIdTransaction() string
	All() []entity.Transaction
	FindByID(transactionID string) entity.Transaction
	Update(b dto.TransactionUpdateDTO) entity.Transaction
	Delete(transactions entity.Transaction)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

//NewTransactionService creates a new instance of AuthService
func NewTransactionService(transactionRep repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRep,
	}
}

func RandomNumbers(n int) string {
	
	const letterBytes = "0123456789"
    b := make([]byte, n)
	// fmt.Println("b=", b)

    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func (service *transactionService) CreateIdTransaction() string {
	rand.Seed(time.Now().UnixNano())
	id_transaction := "TRN-" +RandomNumbers(12) + "-" + RandomLetters(6)
	return id_transaction
}

func (service *transactionService) CreateTransaction(transaction dto.TransactionCreateDTO) entity.Transaction {
	transactionToCreate := entity.Transaction{}
	
	err := smapping.FillStruct(&transactionToCreate, smapping.MapFields(&transaction))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	
	res := service.transactionRepository.InsertTransaction(transactionToCreate)
	return res
}

func (service *transactionService) FindByID(transactionID string) entity.Transaction {
	return service.transactionRepository.FindTransactionId(transactionID)
}

func (service *transactionService) All() []entity.Transaction {
	return service.transactionRepository.AllTransactions()
}

func (service *transactionService) Update(transaction dto.TransactionUpdateDTO) entity.Transaction {
	transactionToUpdate := entity.Transaction{}
	// Fill the variable
	err := smapping.FillStruct(&transactionToUpdate, smapping.MapFields(&transaction))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	// Update the variable
	updatedTransaction := service.transactionRepository.UpdateTransaction(transactionToUpdate)
	return updatedTransaction
}

func (service *transactionService) Delete(transaction entity.Transaction) {
	service.transactionRepository.DeleteTransaction(transaction)
}



