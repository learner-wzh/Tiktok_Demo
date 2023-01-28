package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok_Demo/config"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := DB.DB()

	// 设置数据库连接池
	sqlDB.SetMaxOpenConns(10)
}
