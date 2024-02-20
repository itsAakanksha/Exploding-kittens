package user

import (
	"context"
	//   "encoding/json"
	"errors"
	"fmt"
	"github.com/itsAakanksha/Exploding-kittens/backend/cache"
	"github.com/redis/go-redis/v9"
	// "strconv"
	"strings"
)

// User represents a user with username and wins
type User struct {
	Username string `json:"username"`
	Wins     int    `json:"wins"`
}

// Error types for specific error handling
var (
	ErrUserNotFound = errors.New("user not found")
)

// CreateUser creates a new user in Redis
func CreateUser(ctx context.Context, client *cache.Client, username string) error {
	key := fmt.Sprintf("user:%s", username)
	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if exists == 1 {
		return errors.New("user already exists")
	}

	return client.Set(ctx, key, 0, 0).Err() // Store wins as an integer
}

// GetUser retrieves a user from Redis
func GetUser(ctx context.Context, client *cache.Client, username string) (*User, error) {
	key := fmt.Sprintf("user:%s", username)
	wins, err := client.Get(ctx, key).Int64()
	if err == redis.Nil {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return &User{Username: username, Wins: int(wins)}, nil
}

// UpdateUserWins updates the user's wins using INCR
func UpdateUserWins(ctx context.Context, client *cache.Client, username string) error {
	key := fmt.Sprintf("user:%s", username)
	result, err := client.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	fmt.Printf("User %s wins updated to %d\n", username, result)
	return nil
}

func GetAllUsersWins(ctx context.Context, client *cache.Client) ([]User, error) {
	// Get all keys matching the user pattern
	userKeys, err := client.Keys(ctx, "user:*").Result()
  fmt.Println("user ::",userKeys)
	if err != nil {
		return nil, err
	}

	// Initialize a slice to store users and their wins
	var users []User

	// Iterate over user keys
	for _, userKey := range userKeys {
		// Extract username from the key
		username := strings.TrimPrefix(userKey, "user:")

		// Retrieve wins for the user
		wins, err := client.Get(ctx, userKey).Int64()
		if err == redis.Nil {
			// User not found, handle as needed
			fmt.Printf("Warning: User %s not found\n", username)
			continue
		} else if err != nil {
			// Handle other errors
			return nil, err
		}

		// Append user to the slice
		users = append(users, User{Username: username, Wins: int(wins)})
	}

	return users, nil
}