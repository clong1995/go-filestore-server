package mq

import (
	"github.com/streadway/amqp"
	"go-filestore-server/config"
	"log"
)

// 将要写到rabbitmq的数据的结构体
type TransferData struct {
	FileHash      string `json:"file_hash"`
	CurLocation   string `json:"cur_location"`
	DestLocation  string `json:"dest_location"`
	DestStoreType int    `json:"dest_store_type"`
}

var done chan bool

// 消费者接收消息
func StartConsume(qName, cName string, callback func(msg []byte) bool) {
	msgs, err := channel.Consume(
		qName,
		cName,
		true,  // 自动应答
		false, // 非唯一的消费者
		false, // rabbitMQ只能设置为false
		false, // noWait, false表示会阻塞直到有消息过来
		nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	done = make(chan bool)

	go func() {
		// 循环读取channel的数据
		for d := range msgs {
			processErr := callback(d.Body)
			if processErr {
				// TODO:将任务写入错误队列，等待后续处理
			}
		}
	}()

	// 接收done的信号，没有信息过来则会一直阻塞，避免该函数退出
	<-done

	// 关闭通道
	channel.Close()
}

// 停止监听队列
func StopConsume() {
	done <- true
}

var conn *amqp.Connection
var channel *amqp.Channel

// 如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

func InitMq() {
	// 是否开启异步转移功能，开启时才初始化rabbitMQ连接
	if !config.DefaultConfig.AsyncTransferEnable {
		return
	}
	if initChannel() {
		channel.NotifyClose(notifyClose)
	}

	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channel = nil
				log.Printf("onNotifyChannelClosed:%+v\n", msg)
				initChannel()
			}
		}
	}()
}

func initChannel() bool {
	if channel != nil {
		return true
	}

	conn, err := amqp.Dial(config.DefaultConfig.RabbitURL)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	if !initChannel() {
		return false
	}

	if nil == channel.Publish(
		exchange,
		routingKey,
		false, // 如果没有对应的queue，就会丢弃这条消息
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		}) {
		return true
	}
	return false
}
