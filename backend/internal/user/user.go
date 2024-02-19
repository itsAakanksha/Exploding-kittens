package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// User struct represents a user with username and wins
type User struct {
	Username string `json:"username"`
	Wins     int    `json:"wins"`
}

// CreateUser creates a new user in the Redis database
func CreateUser(ctx context.Context, client *redis.Client, username string) error {
	user := User{Username: username, Wins: 0}
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return client.Set(ctx, fmt.Sprintf("user:%s", username), userJSON, 0).Err()
}

// GetUser retrieves a user's information from the Redis database
func GetUser(ctx context.Context, client *redis.Client, username string) (*User, error) {
	userJSON, err := client.Get(ctx, fmt.Sprintf("user:%s", username)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err // Handle other potential errors
	}
	var user User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserWins updates the user's win count in the Redis database
func UpdateUserWins(ctx context.Context, client *redis.Client, username string) error {
	user, err := GetUser(ctx, client, username)
	if err != nil {
		return err // Handle any error from GetUser
	}
	user.Wins++
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return client.Set(ctx, fmt.Sprintf("user:%s", username), userJSON, 0).Err()
}
