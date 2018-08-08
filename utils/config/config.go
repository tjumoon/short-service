package config

import (
	"time"
	"log"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	// App
	RunMode string

	// Server
	HttpPort 		int
	ReadTimeout 	time.Duration
	WriteTimeout 	time.Duration

	// Redis
	RedisHost		string
	RedisMaxIdle	int
	RedisMaxActive	int
	RedisIdleTimeout time.Duration
	RedisDB			string
	RedisPassword   string


	// log
	AccessLogFilePath      string      = "log/access"
	AccessLogFileExtension string   = ".txt"
	AccessLogMaxSize       int    = 5 // megabytes
	AccessLogMaxBackups    int    = 7
	AccessLogMaxAge        int    = 30 //days
	ErrorLogFilePath       string   = "log/error"
	ErrorLogFileExtension  string    = ".json"
	ErrorLogMaxSize        int    = 10 // megabytes
	ErrorLogMaxBackups     int    = 7
	ErrorLogMaxAge         int    = 30 //days

)

func init()  {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	loadBase()
	loadServer()
	loadRedis()
	loadLog()
}

func loadBase()  {
	RunMode = Cfg.Section("").Key("RUN_Mode").MustString("debug")
}

func loadServer()  {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("Write_TIMEOUT").MustInt(60)) * time.Second
}

func loadRedis()  {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis': %v", err)
	}
	RedisHost = sec.Key("HOST").MustString("127.0.0.1:6379")
	RedisMaxIdle = sec.Key("MAX_IDLE").MustInt(1)
	RedisMaxActive = sec.Key("MAX_ACTIVE").MustInt(10)
	RedisIdleTimeout = time.Duration(sec.Key("IDLE_TIMEOUT").MustInt(180)) * time.Second
	RedisDB = sec.Key("DB").MustString("myredis")
	RedisPassword = sec.Key("PASSWORD").MustString("3870432")
}

func loadLog()  {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		log.Fatalf("Fail to get section 'log': %v", err)
	}

	AccessLogFilePath = sec.Key("ACCESS_LOG_FILE_PATH").MustString("log/accesslog")
	AccessLogFileExtension = sec.Key("ACCESS_LOG_FILE_EXTENSION").MustString(".txt")
	AccessLogMaxSize = sec.Key("ACCESS_LOG_MAX_SIZE").MustInt(5)
	AccessLogMaxBackups = sec.Key("ACCESS_LOG_MAX_BACKUPS").MustInt(7)
	AccessLogMaxAge = sec.Key("ACCESS_LOG_MAX_AGE").MustInt(30)
	ErrorLogFilePath = sec.Key("ERROR_LOG_FILE_PATH").MustString("log/errorlog")
	ErrorLogFileExtension = sec.Key("ERROR_LOG_FILE_EXTENSION").MustString(".json")
	ErrorLogMaxSize = sec.Key("ERROR_LOG_MAX_SIZE").MustInt(10)
	ErrorLogMaxBackups = sec.Key("ERROR_LOG_MAX_BACKUPS").MustInt(7)
	ErrorLogMaxAge = sec.Key("ERROR_LOG_MAX_AGE").MustInt(30)
}