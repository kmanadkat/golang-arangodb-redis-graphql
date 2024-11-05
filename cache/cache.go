package cache

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var CacheClient *redis.Client

func InitializeCache() {
	var addr string = os.Getenv("CACHE_HOST") + ":" + os.Getenv("CACHE_PORT")
	db, err := strconv.Atoi(os.Getenv("CACHE_DB"))
	if err != nil {
		fmt.Println("Error converting CACHE_DB to int: ", err)
		panic(err)
	}

	CacheClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // No password set
		DB:       db, // Use default DB
		Protocol: 2,  // Connection protocol
	})
}
