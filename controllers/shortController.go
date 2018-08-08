package controllers

import (
	"github.com/gin-gonic/gin"
	"short-service/utils/crypto"
	"short-service/controllers/response"
	"short-service/utils/e"
	"short-service/utils/lru"
	"short-service/utils/count"
	"log"
	"short-service/utils/redis"
)

const URL_PREFIX = "http://simonblog.cn/"

const MAXMUM = 1000000

var (
	lruCache *lru.LRU
)

func init(){
	lruCache = lru.NewLRUCache(MAXMUM)
}

func Short(c *gin.Context)  {
	originUrl := c.Query("originUrl")
	log.Printf("originUrl %v\n", originUrl)
	if len(originUrl) == 0 {
		response.E(c, e.INVALID_PARAMS)
		return
	}

	key := crypto.GetMD5(originUrl)
	var shortUrl string
	// LRUCache存在
	if lruCache.IsExist(key) {
		shortUrl = lruCache.Get(key).(string)
	} else {
		r, err := redis.GetString(key)
		// LRUCache不存在，redis存在
		if err == nil {
			lruCache.Add(key, r)
			shortUrl = r
		} else {
			shortUrl = count.Generate()
			lruCache.Add(key, shortUrl)
			lruCache.Add(shortUrl, originUrl)
			redis.Set(key, shortUrl)
			redis.Set(shortUrl, originUrl)
		}

	}
	log.Printf("shortUrl: %v\n", shortUrl)
	response.J(c, gin.H{
		"shortUrl": shortUrl,
	})
}

func Expand(c *gin.Context)  {
	shortUrl := c.Query("shortUrl")
	log.Printf("shortUrl: %v\n", shortUrl)
	if len(shortUrl) == 0 {
		response.E(c, e.INVALID_PARAMS)
		return
	}
	var originUrl string
	// LRUCache存在
	if lruCache.IsExist(shortUrl) {
		originUrl = lruCache.Get(shortUrl).(string)
	} else {
		r, err := redis.GetString(shortUrl)
		if err != nil {
			response.E(c, e.URL_NOT_EXIST)
			return
		}
		originUrl = r
		lruCache.Add(shortUrl, originUrl)
	}
	log.Printf("originUrl %v\n", originUrl)
	response.J(c, gin.H{
		"originUrl": originUrl,
	})

}

