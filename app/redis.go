package app

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mohammaderm/todoList/config"
	"github.com/mohammaderm/todoList/log"
)

func ConnectRedis(logger log.Logger, config *config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Server + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		logger.Panic("cant not connect to redis", map[string]interface{}{
			"error": err.Error(),
		})
	}
	fmt.Println(pong)
	fmt.Println("redis is connected succesfully")

	return client
}
