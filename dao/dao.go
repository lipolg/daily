package dao

import (
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

type Result struct {
	PG       string `json:"pg"`
	AppMode  string `json:"app_mode"`
	HttpPort string `json:"http_port"`
}

var (
	Res       Result
	NacosAddr string
	NacosPort int
	DB        *gorm.DB
)

func InitDao() {
	file, err := ini.Load("./dao/conf.ini")
	if err != nil {
		panic(err)
	}
	loadNacos(file)
	linkNacos()
	linkPg()
}
