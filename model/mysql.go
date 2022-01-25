package model

import (
	"VideoWeb/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func InitMysql() {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.MysqlUSER,
			config.MysqlPW,
			config.MysqlADDR,
			config.MysqlDB,
		))
	db.LogMode(true)
	if err != nil {
		fmt.Println("error mysql")
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
	Migrate()
}
func Migrate() {
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Video{}).AutoMigrate(&Comment{})
}
