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

//LenderService is a contract about something that this service can do
type LenderService interface {
	CreateLender(lender dto.LenderCreateDTO) entity.Lender
	CreateIdLender() string
	All() []entity.Lender
	FindByID(lenderID string) entity.Lender
	Update(b dto.LenderUpdateDTO) entity.Lender
	Delete(lenders entity.Lender)
}

type lenderService struct {
	lenderRepository repository.LenderRepository
}

//NewLenderService creates a new instance of AuthService
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
	rand.Seed(time.Now().UnixNano())
	id_lender := "LDR-" + RandomLetters(6) + "-" + GetCurrentTime()
	return id_lender
}

func (service *lenderService) CreateLender(lender dto.LenderCreateDTO) entity.Lender {
	lenderToCreate := entity.Lender{}
	
	err := smapping.FillStruct(&lenderToCreate, smapping.MapFields(&lender))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	
	res := service.lenderRepository.InsertLender(lenderToCreate)
	return res
}

func (service *lenderService) FindByID(lenderID string) entity.Lender {
	return service.lenderRepository.FindLenderId(lenderID)
}

func (service *lenderService) All() []entity.Lender {
	return service.lenderRepository.AllLenders()
}

func (service *lenderService) Update(lender dto.LenderUpdateDTO) entity.Lender {
	lenderToUpdate := entity.Lender{}
	// Fill the variable
	err := smapping.FillStruct(&lenderToUpdate, smapping.MapFields(&lender))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	// Update the variable
	updatedLender := service.lenderRepository.UpdateLender(lenderToUpdate)
	return updatedLender
}

func (service *lenderService) Delete(lender entity.Lender) {
	service.lenderRepository.DeleteLender(lender)
}



