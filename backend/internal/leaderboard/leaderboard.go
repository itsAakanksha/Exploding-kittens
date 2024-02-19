package leaderboard

import (
  "context"
//   "encoding/json"
  "fmt"
  "sort"

  "github.com/redis/go-redis/v9"
  "github.com/itsAakanksha/Exploding-kittens/backend/internal/user" // Adjust the import path as needed
)

const (
  leaderboardKey = "leaderboard"
  maxUsers      = 10 // Number of users to display on the leaderboard
)

// GetTopNUsers retrieves the top N users from the leaderboard
func GetTopNUsers(ctx context.Context, client *redis.Client, numUsers int) ([]user.User, error) {
  if numUsers <= 0 {
    return nil, fmt.Errorf("invalid number of users: %d", numUsers)
  }

  // Ensure requested number of users doesn't exceed maximum
  numUsers = min(numUsers, maxUsers)

  // Use ZRANGE and ZSCORE to get top users and their scores separately
  users, err := client.ZRange(ctx, leaderboardKey, 0, -1).Result()
  if err != nil {
    return nil, err
  }

  scores := make([]float64, len(users))
  for i, username := range users {
    score, err := client.ZScore(ctx, leaderboardKey, username).Result()
    if err != nil {
      return nil, err
    }
    scores[i] = score
  }

  // Sort users and scores together based on scores in descending order
  type userScorePair struct {
    username string
    score    float64
  }
  userScorePairs := make([]userScorePair, len(users))
  for i, username := range users {
    userScorePairs[i] = userScorePair{username: username, score: scores[i]}
  }
  sort.Slice(userScorePairs, func(i, j int) bool {
    return userScorePairs[i].score > userScorePairs[j].score
  })

  // Extract top N users and convert to User structs
  leaderboard := make([]user.User, min(numUsers, len(userScorePairs)))
  for i := 0; i < min(numUsers, len(userScorePairs)); i++ {
    leaderboard[i] = user.User{
      Username: userScorePairs[i].username,
      Wins:     int(userScorePairs[i].score),
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
