package leaderboard

import (
	"context"
	// "encoding/json"
	"fmt"
	// "sort"

	"github.com/redis/go-redis/v9"
	"github.com/itsAakanksha/Exploding-kittens/backend/internal/user"
)

const (
	leaderboardKey = "leaderboard"
	maxUsers      = 10 // Number of users to display on the leaderboard
)

// GetTopNUsers retrieves the top N users from the leaderboard
func GetTopNUsers(ctx context.Context, client *redis.Client, numUsers int) ([]User, error) {
	if numUsers <= 0 {
		return nil, fmt.Errorf("invalid number of users: %d", numUsers)
	}

	// Ensure requested number of users doesn't exceed maximum
	numUsers = min(numUsers, maxUsers)

	// Use ZRANGE with REV and WITHSCORES to get top users with their scores
	users, err := client.ZRangeWithScores(ctx, leaderboardKey, 0, -1, &redis.ZRangeWithScores{Rev: true}).Result()
	if err != nil {
		return nil, err
	}

	// Convert redis.ZMember to User structs
	leaderboard := make([]User, len(users))
	for i, member := range users {
		leaderboard[i] = User{
			Username: member.Member.(string),
			Wins:     int(member.Score),
		}
	}

	return leaderboard, nil
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
