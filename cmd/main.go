package main

import (
	"context"
	"event_auth/internal/authorization"
	"event_auth/internal/brocker/nats"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	godotenv.Load()
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Помилка підключення до REDIS:", err)
		return
	}

	nc, err := nats.NewNatsClient()
	if err != nil {
		fmt.Println("Помилка підключення до REDIS:", err)
		return
	}
	defer nc.Close()
	//fmt.Println("Підключено до Redis:", pong)

	authModule := authorization.NewAuthorizationModule(rdb, nc)
	authModule.InitNatsSubscribers()

	select {}
}

type AuthorizationModule struct {
	cashe *redis.Client
}
