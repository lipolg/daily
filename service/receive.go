package service

import (
	"daily/dao"
	"daily/model"
	"strings"
)

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// 新增消息
func AddWX(tempStr string) {
	str := strings.Split(tempStr, " ")
	wechat := model.WeChat{
		Label:   str[0],
		Content: str[1],
	}
	err := dao.DB.Create(&wechat).Error
	if err != nil {
		panic(err)
	}

}

// 删除消息
func DeleteWX(tempStr string) {
	err := dao.DB.Where("id = ?", tempStr[1:]).Delete(&model.WeChat{}).Error
	if err != nil {
		panic(err)
	}
}
