package main

import (
	"newsfeeder/httpd/handler"
	"github.com/gin-gonic/gin"
	"newsfeeder/platform/newsfeed"

)
 
func main() {
	feed := newsfeed.New()
	r := gin.Default()

	api := r.Group("/api")
	{	
		api.GET("/ping", handler.PingGet())
		api.GET("/newsfeed", handler.NewsfeedGet(feed))
		api.POST("/newsfeed", handler.NewsfeedPost(feed))
	}
	r.Run("0.0.0.0:5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}