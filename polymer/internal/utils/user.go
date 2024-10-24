package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func FetchUserData(userID int) (*User, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching user data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &user, nil
}
