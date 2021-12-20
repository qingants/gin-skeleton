package xgorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func Open(dsn string, maxIdleConnections, maxOpenConnections int) *gorm.DB {
	db, err := gorm.Open("mysql", dsn)
	zap.S().Infof("Connect Mysql [%v]", dsn)
	if err != nil {
		zap.L().Fatal("Connect Mysql Error ", zap.Error(err), zap.String("dsn", dsn))
		return nil
	}

	db.SingularTable(true)
	db.LogMode(false)
	db.DB().SetMaxIdleConns(maxIdleConnections)
	db.DB().SetMaxOpenConns(maxOpenConnections)

	if err := db.DB().Ping(); err != nil {
		zap.L().Fatal("Open mysql error: %v", zap.Error(err))
		return nil
	}

	return db
}

func Close(db *gorm.DB) {
	if err := db.Close(); err != nil {
		zap.L().Error("Close Mysql Error", zap.Error(err))
	}
}
