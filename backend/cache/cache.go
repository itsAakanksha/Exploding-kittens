package cache

import (
	"context"
	"fmt"
	"log"
   
	"github.com/redis/go-redis/v9"
)


type Client struct {
	*redis.Client
}
 

func New(Addr string,Username string,RedisPassword string,RedisDB int) (*Client, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		// Addr:     "redis-13277.c322.us-east-1-2.ec2.cloud.redislabs.com:13277",
		// Username: "aku",
		// Password: "@12Password", 
		// DB:       0,
		Addr:     Addr,
		Username: Username,
		Password: RedisPassword,
		DB:       RedisDB,
	})
	
	// defer client.Close()
	status, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err)
		// fmt.Println(err)

		
	}
	fmt.Println("k",status)

	return &Client{client}, nil
}
