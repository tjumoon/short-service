package common

import "short-service/utils/redis"

func Startup()  {
	redis.InitRedisPool()
}
