package main

import (
	"fmt"
	"math/rand"

	"github.com/ayushh-vermaa/polymer/internal/rewards"
	"github.com/ayushh-vermaa/polymer/internal/shop"
	"github.com/ayushh-vermaa/polymer/store"
)

func main() {
	cardListPtr, _ := rewards.FetchCardList()

	var cardKeys []string
	for _, cardIssuer := range *cardListPtr {
		for _, card := range cardIssuer.Card {
			cardKeys = append(cardKeys, card.CardKey)
		}
	}

	rand.Shuffle(len(cardKeys), func(i, j int) {
		cardKeys[i], cardKeys[j] = cardKeys[j], cardKeys[i]
	})

	client, err := store.ConnectMongoDB()
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %s", err)
		return
	}

	wallet := shop.BuildWallet(client, cardKeys[:100])

	domainName := "amazon.com"
	amount := 127.56
	_ = shop.Transact(client, domainName, amount, wallet)
}
