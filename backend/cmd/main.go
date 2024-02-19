// cmd/main.go
package main

import (
	"fmt"
)

func main(){
	c:=cache.New()
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
