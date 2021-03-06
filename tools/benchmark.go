package main

import (
	"context"
	"fmt"
	redisbloom "github.com/RedisBloom/redisbloom-go"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

func main() {
	var client = redisbloom.NewClient("localhost:6379", "nohelp", nil)
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	start := time.Now()
	for w := 0; w <= 1000000; w++ {
		client.Exists("MY-BLOOM", "0123456789-099999")
	}
	elapsed := time.Since(start)
	fmt.Println("BF took:", elapsed)
	start = time.Now()
	for w := 0; w <= 1000000; w++ {
		rdb.SIsMember(ctx, "MY-SET", "0123456789-099999")
	}
	elapsed = time.Since(start)
	fmt.Println("Set took:", elapsed)

}
