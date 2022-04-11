package conf

import (
	"daily/dao"
	"daily/model"
)

func InitConf() {
	dao.InitDao()
	model.Migration()
}
