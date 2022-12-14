package routes

import (
	"gin_sample_project/app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(router *gin.Engine) {
	router.GET("/", controllers.HomeHandler)
	router.GET("/hello", controllers.HelloHandler)

	router.GET("/something", controllers.SomethingHandler)

	router.GET("/nested-template", controllers.NestedTemplateHandler)
}
