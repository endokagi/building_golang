package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "167.99.29.46:6379",
		Password: "etneca-th", // no password set
		DB:       0,           // use default DB
	})

	// err := rdb.Set(ctx, "refreshToken", "wasd123456789aadghghjjgkghgggfjffdfhdgh", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	val, err := rdb.Del(ctx, "refreshToken").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("refreshToken", val)
}
