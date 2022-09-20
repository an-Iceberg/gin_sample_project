package routes

import (
	"gin/app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(router *gin.Engine) {
	router.GET("/", controllers.HomeHandler)
	router.GET("/hello", controllers.HelloHandler)
}
