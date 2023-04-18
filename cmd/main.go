package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"jd_workout_golang/app/middleware"
	docs "jd_workout_golang/docs"
	"jd_workout_golang/internal/router"
)

func main() {
	r := SetupRouter()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":80") // listen and serve on
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	// 註冊 router group
	apiGroup := r.Group("/api")
	// 註冊 user router
	router.RegisterUser(apiGroup)

	return r
}
