package cache

import (
  "context"
  "fmt"
  "github.com/redis/go-redis/v9"
)

// Client represents a connection to a Redis server
type Client struct {
  *redis.Client
}

// New creates a new Client instance with the provided configuration
func New(addr string, password string, db int) (*Client, error) {
  client := redis.NewClient(&redis.Options{
    Addr:     addr,
    Password: password,
    DB:       db,
  })

  if err := client.Ping(context.Background()).Err(); err != nil {
    return nil, fmt.Errorf("failed to connect to Rediss: %w", err) // Return specific error type
  }

  return &Client{client}, nil
}

// Ping checks the connection to the Redis server
func (c *Client) Ping(ctx context.Context) error {
  return c.Client.Ping(ctx).Err()
}











// package cache 

// import (
// 	"context"
// 	"github.com/redis/go-redis/v9"
// )

// type Cache struct {
// 	*redis.Client
// }

// func New(addr string , passwd string , db int) *Cache{
// 	client := redis.NewClient(&redis.Options {
// 		Addr: addr,
// 		Password: passwd,
// 		DB: db,
// 	})
// 	return &Cache{client}
// }

// func (c *Cache) Ping(ctx context.Context) error {
// 	return c.Client.Ping(ctx).Err()
// }