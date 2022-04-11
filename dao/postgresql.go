package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func linkPg() {
	db, err := gorm.Open(postgres.Open(Res.PG), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
