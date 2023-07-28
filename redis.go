package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func rdInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDISHOST"), os.Getenv("REDISPORT")),
		Password: os.Getenv("REDISPASSWORD"),
		DB:       0,
	})

	return client
}

func setKey(r *redis.Client, k string, v string) {

	ctx := context.Background()

	err := r.Set(ctx, k, v, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("key set successfully")

}

func getKey(r *redis.Client, k string) string {

	ctx := context.Background()

	val, err := r.Get(ctx, k).Result()
	if err != nil {
		panic(err)
	}

	return val

}
