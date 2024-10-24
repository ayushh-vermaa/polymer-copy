package shop

import (
	"log"

	"github.com/ayushh-vermaa/polymer/internal/rewards"
	"github.com/ayushh-vermaa/polymer/store"

	// For doing exact math with strings
	"go.mongodb.org/mongo-driver/mongo"
)

// BaseWallet represents a collection of cards without personal info.
type BaseWallet struct {
	Cards []*rewards.CardDetail `json:"cards"`
}

// BuildWallet gets cards for a given set of cardKey strings and builds a
// BaseWallet instance with them.
func BuildWallet(client *mongo.Client, cardKeys []string) *BaseWallet {
	var wallet BaseWallet

	cards, err := GetCards(client, cardKeys)
	if err != nil {
		log.Printf("Error getting details for cards: %v", err)
		return &wallet
	}

	wallet.Cards = cards
	return &wallet
}

// SelectBest finds the card with the highest reward value for the given
// merchant category. It returns the reward value as a string and the card with
// the best reward.
func (wallet *BaseWallet) SelectBest(categoryID int) *store.CardDetails {
	bestCardDetails := store.CardDetails{
		CardKey:  "",
		CardName: "",
		RewardDetails: store.RewardDetails{
			Amount:          0.0,
			Currency:        "",
			CashConvertible: false,
			CashConvValue:   0.0,
			Value:           0.0,
		},
	}

	for _, card := range wallet.Cards {
		rewardDetails := CalculateBonusValue(categoryID, card)
		if rewardDetails.Value > bestCardDetails.RewardDetails.Value {
			bestCardDetails = store.CardDetails{
				CardKey:       card.CardKey,
				CardName:      card.CardName,
				RewardDetails: *rewardDetails,
			}
		}
	}

	return &bestCardDetails
}
