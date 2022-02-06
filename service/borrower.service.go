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

//BorrowerService is a contract about something that this service can do
type BorrowerService interface {
	// VerifyCredential(email string, password string) interface{}
	// IsDuplicateEmail(email string) bool
	CreateBorrower(borrower dto.BorrowerCreateDTO) entity.Borrower
	CreateIdBorrower() string
	All() []entity.Borrower
}

type borrowerService struct {
	borrowerRepository repository.BorrowerRepository
}

//NewBorrowerService creates a new instance of AuthService
func NewBorrowerService(borrowerRep repository.BorrowerRepository) BorrowerService {
	return &borrowerService{
		borrowerRepository: borrowerRep,
	}
}

func RandomLettersNumbers(n int) string {
	
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// fmt.Println("b[i]=", letterBytes[rand.Intn(len(letterBytes))])
	// fmt.Println("rand Intn len=", rand.Intn(len(letterBytes)))
	
    b := make([]byte, n)
	// fmt.Println("b=", b)

    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func GetCurrentTime() string {
	currentTime := time.Now().Format("2006-01-02")
	return currentTime
}

func (service *borrowerService) All() []entity.Borrower {
	return service.borrowerRepository.AllBorrowers()
}


func (service *borrowerService) CreateIdBorrower() string {
	rand.Seed(time.Now().UnixNano())
	id_borrower := "BRW-" + RandomLettersNumbers(8) + "-" + GetCurrentTime()
	return id_borrower
}

func (service *borrowerService) CreateBorrower(borrower dto.BorrowerCreateDTO) entity.Borrower {
	borrowerToCreate := entity.Borrower{}
	
	err := smapping.FillStruct(&borrowerToCreate, smapping.MapFields(&borrower))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	
	res := service.borrowerRepository.InsertBorrower(borrowerToCreate)
	return res
}



