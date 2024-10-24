package rewards

import (
	"encoding/json"
)

// CardListResponse represents API response to fetchCardList.
type CardListResponse []struct {
	CardIssuer string `json:"cardIssuer"`
	Card       []struct {
		CardKey  string `json:"cardKey"`
		CardName string `json:"cardName"`
	} `json:"card"`
}

// FetchCardList fetches list of cards from the API in format of
// CardListResponse and returns with any error.
func FetchCardList() (*CardListResponse, error) {
	var params []string
	resp, err := FetchEndpoint("card_list", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cardList CardListResponse
	if err := json.NewDecoder(resp.Body).Decode(&cardList); err != nil {
		return nil, err
	}

	return &cardList, nil
}
