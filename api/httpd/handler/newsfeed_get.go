package handler

import(
	"github.com/gin-gonic/gin"
	"newsfeeder/platform/newsfeed"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc{
	return func(c *gin.Context){
			results := feed.GetAll()
			c.JSON(http.StatusOK,results)
	}
}