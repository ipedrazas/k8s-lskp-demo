package main

import (
    "fmt"
    "gopkg.in/redis.v3"
    "os"
    "time"
)

var redisClient *Redis

type Redis struct {
    *redis.Client
}

func main() {
    uri := fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST"))
    redisClient = &Redis{
        redis.NewClient(&redis.Options{
            Addr:     uri,
            Password: "", // no password set
            DB:       0,  // use default DB
        }),
    }

    count := 1
    for count < 10 {
        redisClient.Incr("ping")
        time.Sleep(30000 * time.Millisecond)
        count ++
    }


}

