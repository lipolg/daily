package service

import (
	"daily/dao"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// WXMsgReply 微信消息回复
func ReplyService(c *gin.Context, fromUser, toUser string) {
	var pushContext []PushContext
	err := dao.DB.Table("we_chats").Select([]string{"id", "content", "label"}).Where("deleted_at is null").Order("id ASC").Scan(&pushContext).Error
	if err != nil {
		logging.Info(err)
	}
	var msg string
	for _, v := range pushContext {
		if len(v.Lable) > 0 {
			msg += "[" + v.Lable + "]"
		}
		msg += v.Content + "[" + strconv.Itoa(int(v.Id)) + "]" + "%0D%0A"
	}
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      msg,
	}

	str, err := xml.Marshal(&repTextMsg)
	if err != nil {
		logging.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(str)
}
