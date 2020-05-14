// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			// set the HTTP status to 200(OK)
			http.StatusOK,
			// use index.html template
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.Run()
}