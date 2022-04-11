package api

import (
	"daily/service"
	"github.com/gin-gonic/gin"
)

// 微信消息回复
func ReplyWX(c *gin.Context, fromUser, toUser string) {
	service.ReplyService(c, fromUser, toUser)
}
