package main

import (
	"daily/api"
	"daily/conf"
	"daily/dao"
	"daily/routes"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	conf.InitConf()
	//每天9点推送
	c := cron.New()
	_, err := c.AddFunc("0 9 * * *", api.PushWX)
	if err != nil {
		log.Println(err)
	}
	c.Start()
	gin.SetMode(dao.Res.AppMode)
	r := routes.NewRouter()
	_ = r.Run(dao.Res.HttpPort)
}
