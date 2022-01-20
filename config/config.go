package config

import (
	"github.com/go-ini/ini"
)

var (
	GinMode string
	GinPort string

	JwtSecret string

	MysqlUSER string
	MysqlPW   string
	MysqlADDR string
	MysqlDB   string

	RedisADDR string
	RedisPW   string
	RedisDB   string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func Init() {
	LoadConfig()
}

func LoadConfig() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		panic(err)
	}

	GinMode = file.Section("GIN").Key("GinMode").String()
	GinPort = file.Section("GIN").Key("GinPort").String()

	JwtSecret = file.Section("JWT").Key("JwtSecret").String()

	MysqlUSER = file.Section("MYSQL").Key("MysqlUSER").String()
	MysqlPW = file.Section("MYSQL").Key("MysqlPW").String()
	MysqlADDR = file.Section("MYSQL").Key("MysqlADDR").String()
	MysqlDB = file.Section("MYSQL").Key("MysqlDB").String()

	RedisADDR = file.Section("REDIS").Key("RedisADDR").String()
	RedisPW = file.Section("REDIS").Key("RedisPW").String()
	RedisDB = file.Section("REDIS").Key("RedisDB").String()

	AccessKey = file.Section("QINIU").Key("AccessKey").String()
	SecretKey = file.Section("QINIU").Key("SecretKey").String()
	Bucket = file.Section("QINIU").Key("Bucket").String()
	QiniuServer = file.Section("QINIU").Key("QiniuServer").String()
}
