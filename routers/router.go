package routers

import (
	"github.com/gin-gonic/gin"
	"short-service/utils/config"
	"short-service/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	switch config.RunMode {
	case "debug":
		r.Use(gin.Logger())
	case "release":
	}

	r.GET("/shorten", controllers.Short)
	r.GET("/expand", controllers.Expand)


	gin.SetMode(config.RunMode)
	r.Use(gin.Recovery())

	return r
}
