package main

import (
	"context"
    "github.com/rs/cors"
	"encoding/json"
	"fmt"
	"os"
	"github.com/gorilla/mux"
	"github.com/itsAakanksha/Exploding-kittens/backend/cache"
	"github.com/itsAakanksha/Exploding-kittens/backend/internal/user"
	"github.com/joho/godotenv"
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}

func handleUpdateUserWins(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	// Addr    string = ""
	// Username string = ""
	// RedisPassword string = ""
	RedisDB      int    = 0
)

func main() {
	godotenv.Load()
	
	ctx := context.Background()
  
	var err error
	client, err = cache.New(os.Getenv("ADDR"),os.Getenv("USERNAME"),os.Getenv("REDIS_PASSWORD"),RedisDB)
	if err != nil {
		log.Panicf("failed to connect to Redis: %v", err)
	}

	log.Println("connected to Redis")

	r := mux.NewRouter()
  corsHandler := cors.Default().Handler(r)

	r.HandleFunc("/createuser", handleCreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{username}", handleGetUser).Methods(http.MethodGet)
  r.HandleFunc("/users/{username}/wins", handleUpdateUserWins).Methods(http.MethodPut)
  r.HandleFunc("/leaderboard", handleGetAllUsersWins).Methods(http.MethodGet)
	srv := &http.Server{Addr: ":10000", Handler: corsHandler}
	go func() {
		log.Println("Server started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to listen and serve: %v", err)
		}
	}()

	
	<-ctx.Done()
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Error shutting down server:", err)
	}
	log.Println("Server stopped")
}





