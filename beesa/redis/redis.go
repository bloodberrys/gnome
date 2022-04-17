package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Open connection to Redis
var redisClient = redis.NewClient(&redis.Options{
	Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
	Password: "",
	DB:       0,
})

func Ping() error {
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)

	return nil
}

func rSet(key string, value interface{}) error {
	if len(key) == 0 {
		key = "name"
		value = "alfian"
	}
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func hSet(key string, field string, value interface{}) error {

	err := redisClient.HSet(ctx, key, field, value).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

// Data integrity, set it with multiple pipe and exec.
func rSetPipe() error {
	pipe := redisClient.TxPipeline()
	pipe.Set(ctx, "name", "golang", 0)
	pipe.Set(ctx, "name1", "golang1", 0)
	pipe.Set(ctx, "year", 2009, 0)
	results, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Print(results)
	return nil
}

// Redis Get with key doesn't exist handler
func rGet(key string) error {
	val, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("%v does not exist", key)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("%v: %v", key, val)
	}

	return nil
}

func rGetAllKeys() []string {

	// the key is regex
	key := "name*"
	keys := []string{}

	iter := redisClient.Scan(ctx, 0, key, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}

// set the json value
func RSetComposite(Link interface{}) error {

	json, err := json.Marshal(Link)
	if err != nil {
		panic(err)
	}
	rSet("id1", json)

	return nil
}

func HSetComposite(key string, field string, Link interface{}) error {

	json, err := json.Marshal(Link)
	if err != nil {
		panic(err)
	}

	hSet(key, field, json)

	return nil

}
