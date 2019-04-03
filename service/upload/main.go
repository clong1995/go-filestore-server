package main

import (
	"fmt"
	"go-filestore-server/config"
	"go-filestore-server/db/mq"
	"go-filestore-server/db/mysql"
	"go-filestore-server/db/redis"
	"go-filestore-server/handler"
	"go-filestore-server/logger"
	"go-filestore-server/route"
	"log"
	"os"
	"strings"
)

func main() {
	// 初始化配置
	fmt.Println("初始化配置...")
	config.InitConfig("./config.json")
	fmt.Printf("config:\t%+v\n", config.DefaultConfig)
	
	// 日志配置
	fmt.Println("日志配置...")
	path := "logs"
	mode := os.ModePerm
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
	file, _ := os.Create(strings.Join([]string{path, "log.txt"}, "/"))
	defer file.Close()
	loger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetDefault(loger)
	
	// 数据库初始化
	redis.InitRedis()
	mysql.InitMysql()
	mq.InitMq()
	
	// 目录初始化
	handler.InitTemp()
	
	fmt.Println("初始化路由...")
	// gin framework
	router := route.Router()
	router.Run(config.DefaultConfig.UploadServiceHost)
}
