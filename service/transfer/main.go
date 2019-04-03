package main

import (
	"bufio"
	"encoding/json"
	"go-filestore-server/config"
	"go-filestore-server/db/mq"
	"go-filestore-server/db/oss"
	"go-filestore-server/model"
	"log"
	"os"
)

func main() {
	if !config.DefaultConfig.AsyncTransferEnable {
		log.Println("异步转移文件功能目前被禁用,请检查相关配置")
		return
	}
	log.Println("文件转移服务启动中，开始监听转移任务队列...")
	mq.StartConsume(config.DefaultConfig.TransOSSQueueName, "transfer_oss", ProcessTransfer)
}

// 处理文件转移
func ProcessTransfer(msg []byte) bool {
	log.Println(string(msg))

	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	fin, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	err = oss.Bucket().PutObject(pubData.DestLocation, bufio.NewReader(fin))
	if err != nil {
		log.Println(err.Error())
		return false
	}

	_ = model.UpdateFileLocation(pubData.FileHash, pubData.DestLocation)
	return true
}
