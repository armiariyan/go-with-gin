package repository

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/gorm"
)

//TransactionRepository is contract what TransactionRepository can do to db
type TransactionRepository interface {
	InsertTransaction(transaction entity.Transaction) entity.Transaction
	AllTransactions() []entity.Transaction
	UpdateTransaction(transaction entity.Transaction) entity.Transaction
	FindTransactionId(transactionID string) entity.Transaction
	DeleteTransaction(transaction entity.Transaction)
}

type transactionConnection struct {
	connection *gorm.DB
}

//NewTransactionRepository is creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: db,
	}
}

func (db *transactionConnection) AllTransactions() []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Find(&transactions)
	return transactions
}

func (db *transactionConnection) InsertTransaction(transaction entity.Transaction) entity.Transaction {
	db.connection.Create(&transaction)
	db.connection.Find(&transaction)
	return transaction
}

func (db *transactionConnection) FindTransactionId(transactionID string) entity.Transaction {
	var transaction entity.Transaction
	db.connection.First(&transaction, "id_transaction = ?", transactionID)
	return transaction
}

func (db *transactionConnection) UpdateTransaction(transaction entity.Transaction) entity.Transaction {
	db.connection.Updates(&transaction)
	db.connection.Find(&transaction)
	return transaction
}

func (db *transactionConnection) DeleteTransaction(transaction entity.Transaction) {
	db.connection.Delete(&transaction)
}