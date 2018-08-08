package redis

import (
	"short-service/utils/config"
	"log"
	"time"
	"github.com/gomodule/redigo/redis"
)

var (
	RedisClient *redis.Pool
)

func InitRedisPool()  {
	RedisClient = &redis.Pool {
		MaxIdle: config.RedisMaxIdle,
		MaxActive: config.RedisMaxActive,
		IdleTimeout: config.RedisIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisHost)
			if err != nil {
				log.Fatalf("Redis init fail %v", err)
			}
			//if _, err := c.Do("AUTH", config.RedisPassword); err != nil {
			//	c.Close()
			//	log.Fatalf("Redis Auth fail %v", err)
			//}
			c.Do("SELECT", config.RedisDB)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func Set(args ...interface{}) (reply interface{}, err error)  {
	rc := RedisClient.Get()
	defer rc.Close()
	reply, err = rc.Do("SET", args...)
	return
}

func GetInt(key interface{}) (result int, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	return redis.Int(rc.Do("GET", key))
}

func GetBool(key interface{}) (result bool, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	return redis.Bool(rc.Do("GET", key))
}

func GetString(key interface{}) (result string, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	return redis.String(rc.Do("GET", key))
}

func Delete(key interface{}) (reply interface{}, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	return rc.Do("DEL", key)
}
