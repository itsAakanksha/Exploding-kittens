// cmd/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/itsAakanksha/Exploding-kittens/backend/redis"
	"github.com/itsAakanksha/Exploding-kittens/backend/internal/user"
	"github.com/itsAakanksha/Exploding-kittens/backend/internal/leaderboard"
	"log"
	"net/http"
	// "./internal/leaderboard"
	// "./internal/redis"
	// "./internal/user"
)


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
	RedisAddr     string = "localhost:6379"
	RedisPassword string = ""
	RedisDB       int    = 0
)

func main() {
	ctx := context.Background()
	c := cache.New(RedisAddr, RedisPassword, RedisDB)
	if err := c.Ping(ctx); err != nil {
		log.Panic("failed to connect to Redis")
	}

	log.Println("connected to Redis")
}

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/mux"
// 	"github.com/your_module_name/internal/redis"
// 	"github.com/your_module_name/internal/user"
// 	"github.com/your_module_name/handlers"
// )

// func main() {
// 	// Initialize Redis
// 	if err := redis.Init(); err != nil {
// 		log.Fatalf("Error initializing Redis: %v", err)
// 		os.Exit(1)
// 	}

// 	// Create a new router using Gorilla Mux
// 	r := mux.NewRouter()

// 	// Register your HTTP handlers
// 	r.HandleFunc("/create-user/{username}", handlers.CreateUserHandler).Methods("POST")

// 	// Start the HTTP server
// 	port := 8080
// 	addr := fmt.Sprintf(":%d", port)
// 	fmt.Printf("Server listening on %s\n", addr)
// 	log.Fatal(http.ListenAndServe(addr, r))
// }
