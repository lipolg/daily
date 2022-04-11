package model

import "daily/dao"

func Migration() {
	err := dao.DB.
		AutoMigrate(
			&WeChat{},
		)
	if err != nil {
		panic(err)
	}
}
