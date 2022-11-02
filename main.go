package main

import (
	"gin_sample_project/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("public/views/**/*.html")

	router.Static("/css", "./public/css")
	router.Static("/js", "./public/js")
	router.Static("/fonts", "./public/fonts")
	router.Static("/images", "./public/images")

	routes.RegisterAppRoutes(router)

	router.Run("gin.test:80")
}
