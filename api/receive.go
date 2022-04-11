package api

import (
	"daily/service"
	"github.com/gin-gonic/gin"
	"log"
)

func ReceiveWX(c *gin.Context) {
	var textMsg service.WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}
	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)
	tempStr := textMsg.Content
	if tempStr[:1] == "/" && len(tempStr) > 1 {
		service.DeleteWX(tempStr)
	} else if tempStr[:1] != "/" {
		service.AddWX(tempStr)
	}
	ReplyWX(c, textMsg.ToUserName, textMsg.FromUserName)
}
