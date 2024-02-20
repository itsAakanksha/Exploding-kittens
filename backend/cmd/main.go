package main

import (
	"context"
  "github.com/rs/cors"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/itsAakanksha/Exploding-kittens/backend/cache"
	"github.com/itsAakanksha/Exploding-kittens/backend/internal/user"
	"log"
	"net/http"
)

// Global Redis client with context awareness
var client *cache.Client

// Handler functions
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var userr user.User
	err := json.NewDecoder(r.Body).Decode(&userr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}

	err = user.CreateUser(ctx, client, userr.Username)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error creating user: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	if err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := user.GetUser(ctx, client, username)
	if err != nil {
	handleError(w, http.StatusNotFound, "User not found: %v", username)
		
  handleError(w, http.StatusInternalServerError, "Error retrieving user: %v", err)
  return
}



	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func handleUpdateUserWins(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	username := vars["username"]
  println(username)
	err := user.UpdateUserWins(ctx, client, username)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error updating user wins: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "User wins updated successfully"})
	if err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
	
func handleGetAllUsersWins(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Retrieve all usernames and wins from the sorted map
	users, err := user.GetAllUsersWins(ctx, client)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error retreving user wins: %v", err)
		return
	}

	// Encode data as JSON and write to response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func handleError(w http.ResponseWriter, code int, format string, args ...interface{}) {
	w.WriteHeader(code)
	fmt.Fprintf(w, format, args...)
}

const (
	RedisAddr    string = "localhost:6379"
	RedisPassword string = ""
	RedisDB      int    = 0
)

func main() {
	ctx := context.Background()
  
	// Create Redis client with context and error handling
	var err error
	client, err = cache.New(RedisAddr, RedisPassword, RedisDB)
	if err != nil {
		log.Panicf("failed to connect to Redis: %v", err)
	}

	log.Println("connected to Redis")

	// Create a new router
	r := mux.NewRouter()
  corsHandler := cors.Default().Handler(r)

	// Define routes and handlers
	r.HandleFunc("/createuser", handleCreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{username}", handleGetUser).Methods(http.MethodGet)
  r.HandleFunc("/users/{username}/wins", handleUpdateUserWins).Methods(http.MethodPut)
  r.HandleFunc("/leaderboard", handleGetAllUsersWins).Methods(http.MethodGet)
	// Start HTTP server with graceful shutdown
	srv := &http.Server{Addr: ":8080", Handler: corsHandler}
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
