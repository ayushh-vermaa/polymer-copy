package store

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const TransactionCollection = "transaction"

type BaseTransaction struct {
	TransactionAt   time.Time       `bson:"transaction_at"`
	SpendAmount     float64         `bson:"spend_amount"`
	MerchantDetails MerchantDetails `bson:"merchant_details"`
	CardDetails     CardDetails     `bson:"card_details"`
}

type MerchantDetails struct {
	DomainName   string `bson:"name"`
	CategoryID   int    `bson:"category_id"`
	CategoryName string `bson:"category_name"`
}

type CardDetails struct {
	CardKey       string        `bson:"card_key"`
	CardName      string        `bson:"card_name"`
	RewardDetails RewardDetails `bson:"reward_details"`
}

type RewardDetails struct {
	Amount          float64 `bson:"amount"`
	Currency        string  `bson:"currency"`
	CashConvertible bool    `bson:"cash_convertible"`
	CashConvValue   float64 `bson:"cash_conv_value"`
	Value           float64 `bson:"value"`
}

// Transaction represents the structure of a transaction document in MongoDB.
type Transaction struct {
	*BaseDocument    `bson:",inline"`
	*BaseTransaction `bson:",inline"`
}

// CreateTransaction creates a Transaction document from the given
// baseTransaction.
func CreateTransaction(baseTransation *BaseTransaction) Transaction {
	transaction := Transaction{
		BaseDocument:    &BaseDocument{},
		BaseTransaction: baseTransation,
	}
	transaction.SetID()
	return transaction
}

// InsertTransaction inserts a new Transaction document into the MongoDB
// collection.
func InsertTransaction(client *mongo.Client, baseTransaction *BaseTransaction) (
	*mongo.InsertOneResult, error) {

	domain := CreateTransaction(baseTransaction)
	store := GetStore(client, TransactionCollection)
	return store.InsertDocument(domain)
}
