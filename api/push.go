package api

import (
	"daily/service"
	logging "github.com/sirupsen/logrus"
)

//推送消息
func PushWX() {
	err := service.PushService()
	if err != nil {
		logging.Info(err)
	}

}
