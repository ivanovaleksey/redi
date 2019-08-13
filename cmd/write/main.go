package main

import (
	"math/rand"

	"github.com/go-redis/redis/v7"
)

func main() {
	localRedis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   8,
	})

	const total = 1000000
	for i := 0; i < total; i++ {
		localRedis.SAdd("item:all", RandStringRunes(14))
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
