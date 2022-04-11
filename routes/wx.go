package routes

import (
	"daily/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	w := r.Group("/wx")
	{
		w.GET("/verify", api.Verify)
		w.POST("/receive", api.ReceiveWX)
	}
	return r
}
