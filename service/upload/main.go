package main

import (
	"fmt"
	"go-filestore-server/config"
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
	fmt.Println("logger init...")
	path := "logs"
	mode := os.ModePerm
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
	file, _ := os.Create(strings.Join([]string{path, "log.txt"}, "/"))
	defer file.Close()
	loger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetDefault(loger)
	
	fmt.Println("初始化路由...")
	// gin framework
	router := route.Router()
	router.Run(config.DefaultConfig.UploadServiceHost)
}
