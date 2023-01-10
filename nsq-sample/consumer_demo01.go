package main

import (
	"errors"
	"github.com/nsqio/go-nsq"
	"log"
)

type MyHandler struct {
}

// 实现 Handler 接口
func (h *MyHandler) HandleMessage(message *nsq.Message) error {
	// 业务逻辑
	body := message.Body
	log.Println("msg = ", string(body))
	return errors.New("handler failed")
}

func main() {

	config := nsq.NewConfig()
	// 创建消费者
	consumer, err := nsq.NewConsumer("order", "ch01", config)
	if err != nil {
		log.Fatalln(err)
	}
	// 需要在 ConnectToNSQD 之前添加 Handler
	// 添加消息处理器，启动 goroutine 执行 HandleMessage 方法
	//consumer.AddHandler(&MyHandler{})
	consumer.AddConcurrentHandlers(&MyHandler{}, 2)

	// 连接到nsqd
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatalln(err)
	}

	// 阻塞
	closeChan := make(chan bool, 1)
	<-closeChan

	// 关闭消费者
	consumer.Stop()
}
