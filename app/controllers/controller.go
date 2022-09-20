package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "home", gin.H{
		"css":       "index",
		"pageTitle": "Home Page",
		"title":     "Go Gin",
		"subTitle":  "You are now drinking Gin. Don't drink too much.",
	})
}

func HelloHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "hello", gin.H{
		"css":       "hello",
		"pageTitle": "Hello Page",
		"title":     "Hello World :)",
		"text":      "Hello World, how are you doing tonight?",
	})
}
