package main

import (
	"github.com/gin-gonic/gin"
	"newsfeeder/platform/newsfeed"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGet)
	r.GET("/newsfeed", handler.NewsfeedGet)
	r.GET("/newsfeed", handler.NewsfeedPost)
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}