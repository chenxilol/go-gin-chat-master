package core

import (
	"fmt"
	"go-gin-chat/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("未配置mysql，取消gorm链接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface // 定义Logger接口
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info) // 如果系统变量是debug， 那么将会输出所以日志，包括sql语句
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) //如果系统变量是其他，那么只输出error信息
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ //连接数据库
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] myssql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	log.Println("success Gorm config")
	return db
}
