package repository
//
//import (
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"log"
//	"os"
//	"time"
//)
//
//func init() {
//	con()
//}
//
//var db *gorm.DB
//var err error
//
//func DB() *gorm.DB {
//	s, err2 := db.DB()
//	if err2 == nil {
//		return db
//	}
//	if err2 = s.Ping(); err2 == nil {
//		return db
//	}
//	con()
//	return db
//}
//
//func con() {
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold: time.Second,   // 慢 SQL 阈值
//			LogLevel:      logger.Silent, // Log level
//			Colorful:      false,         // 禁用彩色打印
//		},
//	)
//	//dbc, err = gorm.Open(mysql.Open("staff_yangshiyou:Y7Xx@DcsZqyVR%MiXk=9nFzELrho%7@(rm-bp1uk2jq84jpy9d85.mysql.rds.aliyuncs.com:3306)/shophelpertest"), &gorm.Config{
//	//	Logger: newLogger,
//	//})
//	dsn := "root:123456@(172.16.0.52:3306)/hellogo?charset=utf8&parseTime=True&loc=Local" // DSN data source name
//	db, err = gorm.Open(mysql.New(mysql.Config{
//		DSN:                       dsn,
//		DefaultStringSize:         256,   // string 类型字段的默认长度
//		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
//		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
//		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
//		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
//	}), &gorm.Config{Logger: newLogger})
//	if err != nil {
//		panic(err)
//	}
//	s, _ := db.DB()
//	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
//	s.SetMaxIdleConns(10)
//	// SetMaxOpenConns 设置打开数据库连接的最大数量。
//	s.SetMaxOpenConns(100)
//	// SetConnMaxLifetime 设置了连接可复用的最大时间。
//	s.SetConnMaxLifetime(time.Hour)
//}
