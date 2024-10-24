package shop

import (
	"log"

	"github.com/ayushh-vermaa/polymer/internal/rewards"
	"github.com/ayushh-vermaa/polymer/store"

	"go.mongodb.org/mongo-driver/mongo"
)

// CalculateBonusValue calculates the highest applicable reward value for a
// given merchant category based on the card's rewards and point value.
func CalculateBonusValue(categoryID int,
	card *rewards.CardDetail) *store.RewardDetails {

	// Start with base value
	bestReward := store.RewardDetails{
		Amount:          card.BaseSpendAmount,
		Currency:        card.BaseSpendEarnCurrency,
		CashConvertible: card.BaseSpendEarnIsCash == 1,
		CashConvValue:   card.BaseSpendEarnCashValue,
		Value:           card.RewardValue(card.BaseSpendAmount),
	}

	for _, bonus := range card.SpendBonusCategory {
		if !bonus.IsApplicable(categoryID) {
			continue
		} else if bonus.EarnMultiplier > bestReward.Amount {
			bestReward.Amount = bonus.EarnMultiplier
			bestReward.Value = card.RewardValue(bonus.EarnMultiplier)
		}
	}

	return &bestReward
}

func FetchAndStoreCard(client *mongo.Client,
	cardKey string) *rewards.CardDetail {

	cardDetailPtr, err := rewards.FetchCardDetail(cardKey)
	if err != nil {
		return cardDetailPtr
	}

	var cardDetail rewards.CardDetail
	_, err = store.InsertCard(client, cardDetailPtr)
	if err != nil {
		return &cardDetail
	}

	return &cardDetail
}

// GetCard takes a cardKey string and tries to find the matching CardDetail in
// the database and return it. If not found, it feteches from the API and
// returns the result after storing it in the database.
func GetCards(client *mongo.Client, cardKeys []string) ([]*rewards.CardDetail,
	error) {

	cards, err := store.GetCardsByKeys(client, cardKeys)
	if err != nil {
		return nil, err
	}

	var cardDetails []*rewards.CardDetail
	for cardKey, card := range cards {
		if card == nil {
			cardDetail := FetchAndStoreCard(client, cardKey)
			cardDetails = append(cardDetails, cardDetail)
			log.Printf("Fetched and stored data for: %s", cardDetail.CardName)
		} else {
			cardDetails = append(cardDetails, &card.CardDetail)
			log.Printf("Found data for: %s", card.CardDetail.CardName)
		}
	}

	return cardDetails, nil
}
