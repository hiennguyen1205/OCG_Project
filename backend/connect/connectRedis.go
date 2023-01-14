package connect

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(rdb)
	rdb.Set(context.Background(), "abc", "value-abc", time.Duration(time.Second*199))
}
