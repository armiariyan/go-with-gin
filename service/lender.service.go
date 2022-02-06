package service

import (
	"log"
	"math/rand"

	"github.com/mashingan/smapping"
	"gitlab.com/armiariyan/intern_golang/dto"
	"gitlab.com/armiariyan/intern_golang/entity"
	"gitlab.com/armiariyan/intern_golang/repository"
)

//LenderService is a contract about something that this service can do
type LenderService interface {
	// VerifyCredential(email string, password string) interface{}
	// IsDuplicateEmail(email string) bool
	CreateLender(lender dto.LenderCreateDTO) entity.Lender
	CreateIdLender() string
}

type lenderService struct {
	lenderRepository repository.LenderRepository
}

//NewLenderService creates a new instance of LenderService
func NewLenderService(lenderRep repository.LenderRepository) LenderService {
	return &lenderService{
		lenderRepository: lenderRep,
	}
}

func RandomLetters(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func (service *lenderService) CreateIdLender() string {
	id_lender := "LDR-" + RandomLetters(6) + "-" + GetCurrentTime()
	return id_lender
}

func (service *lenderService) CreateLender(lender dto.LenderCreateDTO) entity.Lender {
	lenderToCreate := entity.Lender{}
	
	err := smapping.FillStruct(&lenderToCreate, smapping.MapFields(&lender))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	// fmt.Println("borrower=", borrower)
	// fmt.Println("borrowerToCreate=", borrowerToCreate)
	
	res := service.lenderRepository.InsertLender(lenderToCreate)
	return res
}


