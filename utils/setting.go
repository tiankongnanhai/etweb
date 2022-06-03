package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	RunMode    string
	ServerIp   string
	ServerPort string

	MysqlIp     string
	MysqlPort   string
	MysqlUser   string
	MysqlPass   string
	MysqlDbName string

	RedisIp   string
	RedisPort string
)

func init() {
	cfg, err := ini.Load("/home/pxr/PXRcode/etweb/config/config.ini")
	if err != nil {
		fmt.Println("config err")
	}
	LoadServer(cfg)
	LoadMysql(cfg)
	LoadRedis(cfg)
}

func LoadServer(cfg *ini.File) {
	RunMode = cfg.Section("").Key("RunMode").String()
	ServerIp = cfg.Section("Server").Key("Ip").String()
	ServerPort = cfg.Section("Server").Key("Port").String()
}

func LoadMysql(cfg *ini.File) {
	MysqlIp = cfg.Section("Mysql").Key("Ip").String()
	MysqlPort = cfg.Section("Mysql").Key("Port").String()
	MysqlUser = cfg.Section("Mysql").Key("User").String()
	MysqlPass = cfg.Section("Mysql").Key("Password").String()
	MysqlDbName = cfg.Section("Mysql").Key("DbName").String()

}

func LoadRedis(cfg *ini.File) {
	RedisIp = cfg.Section("").Key("RunMode").String()
	RedisPort = cfg.Section("Server").Key("Ip").String()

}
