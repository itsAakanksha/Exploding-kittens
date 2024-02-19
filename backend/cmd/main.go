// cmd/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	// "github.com/itsAakanksha/Exploding-kittens/backend/redis"
	"github.com/itsAakanksha/Exploding-kittens/backend/internal/user"
	"github.com/itsAakanksha/Exploding-kittens/backend/internal/leaderboard"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	// "./internal/leaderboard"
	// "./internal/redis"
	// "./internal/user"
)



// Global Redis client with context awareness
var client *redis.Client

// Handler functions (unchanged for brevity)
// ...
// handleCreateUser creates a new user in the database
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	err = user.CreateUser(client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating user: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User created successfully")
}

// handleUpdateUserWins updates the user's win count in the database
func handleUpdateUserWins(w http.ResponseWriter, r *http.Request) {
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  fmt.Fprintf(w, "Invalid request body")
	  return
	}
  
	err = user.UpdateUserWins(client, username)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  fmt.Fprintf(w, "Error updating user wins: %v", err)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User wins updated successfully")
  }

// handleGetLeaderboard retrieves the top N users from the leaderboard
func handleGetLeaderboard(w http.ResponseWriter, r *http.Request) {
	var numUsers int
	_, err := fmt.Fscanf(r.Body, "%d", &numUsers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	users, err := leaderboard.GetTopNUsers(client, numUsers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error retrieving leaderboard: %v", err)
		return
	}

	userJSON, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling user data: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}


const (
  RedisAddr   string = "localhost:6379"
  RedisPassword string = ""
  RedisDB      int = 0
)

func main() {
  ctx := context.Background()

  // Create Redis client with context and error handling
  var err error
  client, err = redis.New(RedisAddr, RedisPassword, RedisDB)
  if err != nil {
    log.Panicf("failed to connect to Redis: %v", err)
  }

  log.Println("connected to Redis")

  // Define HTTP routes and handlers
  http.HandleFunc("/createUser", handleCreateUser)
  http.HandleFunc("/updateUserWins", handleUpdateUserWins)
  http.HandleFunc("/getLeaderboard", handleGetLeaderboard)

  // Start HTTP server with graceful shutdown
  srv := &http.Server{Addr: ":8080"}
  go func() {
    log.Println("Server started on port 8080")
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("failed to listen and serve: %v", err)
    }
  }()

  // Handle graceful shutdown
  <-ctx.Done()
  log.Println("Shutting down server...")
  if err := srv.Shutdown(ctx); err != nil {
    log.Println("Error shutting down server:", err)
  }
  log.Println("Server stopped")
}



// const (
// 	RedisAddr     string = "localhost:6379"
// 	RedisPassword string = ""
// 	RedisDB       int    = 0
// )

// func main() {
// 	ctx := context.Background()
// 	c := cache.New(RedisAddr, RedisPassword, RedisDB)
// 	if err := c.Ping(ctx); err != nil {
// 		log.Panic("failed to connect to Redis")
// 	}
	
// 	log.Println("connected to Redis")
// }
