package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

func main() {
	localRedis := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		DB:          8,
		ReadTimeout: 10000 * time.Millisecond,
		PoolSize:    30,
	})

	pipe := localRedis.Pipeline()

	const size = 50
	cmds := make([]redis.StringSliceCmd, 0, size)

	for i := 0; i < size; i++ {
		cmd := pipe.SMembers("item:all")
		cmds = append(cmds, *cmd)
	}

	_, err := pipe.Exec()
	fmt.Println(err)
}
