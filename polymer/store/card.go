package store

import (
	"context"
	"fmt"
	"time"

	"github.com/ayushh-vermaa/polymer/internal/rewards"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CardCollection = "card"

// Card represents the structure of a credit card document in MongoDB.
type Card struct {
	*BaseDocument `bson:",inline"`
	CardDetail    rewards.CardDetail `bson:"card_detail"`
}

// CreateCard creates a Card document from a given CardDetail object.
func CreateCard(cardDetail *rewards.CardDetail) *Card {
	card := Card{
		BaseDocument: &BaseDocument{},
		CardDetail:   *cardDetail,
	}
	card.SetID()
	return &card
}

// InsertCard inserts a Card document into the cluster from a given CardDetail
// object.
func InsertCard(client *mongo.Client, cardDetail *rewards.CardDetail) (
	*mongo.InsertOneResult, error) {

	card := CreateCard(cardDetail)
	store := GetStore(client, CardCollection)
	return store.InsertDocument(card)
}

// GetCardsByKeys retrieves multiple Card documents based on a slice of unique
// card keys.
func GetCardsByKeys(client *mongo.Client, cardKeys []string) (map[string]*Card,
	error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"card_detail.card_key": bson.M{"$in": cardKeys}}

	store := GetStore(client, CardCollection)
	cursor, err := store.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cards: %w", err)
	}
	defer cursor.Close(ctx)

	cardMap := make(map[string]*Card)
	for cursor.Next(ctx) {
		var card Card
		if err := cursor.Decode(&card); err != nil {
			return nil, fmt.Errorf("failed to decode card: %w", err)
		}
		cardMap[card.CardDetail.CardKey] = &card
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return cardMap, nil
}
