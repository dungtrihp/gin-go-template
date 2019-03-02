package datasources

import (
	"errors"
	"log"
	"start/config"
	"time"

	"github.com/go-redis/redis"
)

type DataCache struct {
	Key            string
	JsonStringData string
	Tll            time.Duration
}

var redisdb *redis.Client

func MyRedisConnect() {
	log.Println("redis connect")
	config := config.GetConfig()
	log.Println(config.GetString("redis.host"), config.GetString("redis.port"), config.GetString("redis.password"))

	redisdb = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.host") + ":" + config.GetString("redis.port"),
		Password: config.GetString("redis.password"),
		DB:       0, // use default DB
	})

	data := redisdb.Ping()
	log.Println(data)
	// pong, err := redisdb.Ping().Result()
	// log.Println("pong: "+pong, "redis error: "+err.Error())
}

// func GetRedis() *redis.Client {
// 	return redisdb
// }

func SetCache(key string, value interface{}, ttl int) error {
	if redisdb == nil {
		return errors.New("redisdb is not connected")
	}

	return redisdb.Set(key, value, time.Duration(ttl*1000000000)).Err()
}

func GetCache(key string) (string, error) {
	if redisdb == nil {
		return "", errors.New("redisdb is not connected")
	}
	return redisdb.Get(key).Result()
}

func IncreaseFlagCounter(key string) (int, error) {
	if redisdb == nil {
		return 0, errors.New("redisdb is not connected")
	}
	result, err := redisdb.Incr(key).Result()
	return int(result), err
}

func DelCacheByKey(keys string) error {
	if redisdb == nil {
		return errors.New("redisdb is not connected")
	}
	redisdb.Del(keys)
	return nil
}
