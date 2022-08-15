package service

import (
	"github.com/go-redis/redis"
	"log"
	"server/model"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     model.Config.Redis.Addr,
		Password: model.Config.Redis.RequirePass,
		DB:       0,
	})
	_, err := Rdb.Ping().Result()
	if err != nil {
		log.Fatal("fail to connect redis:", err)
	}
	log.Printf("Init Redis Successfully")
}
