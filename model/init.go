package model

import (
	"github.com/jinzhu/gorm"
	"github.com/qingants/gin-skeleton/pkg/xgorm"
	"github.com/qingants/gin-skeleton/setting"
)

var (
	rds *gorm.DB
)

func Open() {
	rds = xgorm.Open(setting.RdsDsn, 10, 100)
}

func Close() {
	//models.CloseMainDb()
}

func Get() *gorm.DB {
	return rds
}

type LoveMovie struct {
	ID  uint
	Uid uint
	Mid uint
}


func (LoveMovie) TableName() string {
	return "lover_movie"
}
