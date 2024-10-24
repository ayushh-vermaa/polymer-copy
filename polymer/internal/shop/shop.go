package shop

import (
	"log"
	"time"

	"github.com/ayushh-vermaa/polymer/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type DomainCategory struct {
	ID   int    `json:"categoryID"`
	Name string `json:"categoryName"`
}

// GetDomainCategoryID gets the categoryID for a given domainName from MongoDB.
func GetDomainCategory(client *mongo.Client,
	domainName string) *DomainCategory {

	domain, _ := store.GetDomainByName(client, domainName)
	if domain == nil {
		return &DomainCategory{ID: -1, Name: ""}
	}

	return &DomainCategory{
		ID:   domain.CategoryID,
		Name: domain.CategoryName,
	}
}

func Transact(client *mongo.Client, domainName string, amount float64,
	wallet *BaseWallet) *store.CardDetails {

	category := GetDomainCategory(client, domainName)
	cardDetails := wallet.SelectBest(category.ID)
	log.Printf("Transacting $%.2f with card %q for %.2f%% value back",
		amount, cardDetails.CardName, cardDetails.RewardDetails.Value*100)

	transaction := store.BaseTransaction{
		TransactionAt: time.Now(),
		SpendAmount:   amount,
		MerchantDetails: store.MerchantDetails{
			DomainName:   domainName,
			CategoryID:   category.ID,
			CategoryName: category.Name,
		},
		CardDetails: *cardDetails,
	}

	store.InsertTransaction(client, &transaction)
	return cardDetails
}
