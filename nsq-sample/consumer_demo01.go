package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type MyHandler struct {
}

// 实现 Handler 接口
func (h *MyHandler) HandleMessage(message *nsq.Message) error {
	body := message.Body
	log.Println("msg = ", string(body))
	return nil
}

func main() {

	config := nsq.NewConfig()
	// 创建消费者
	consumer, err := nsq.NewConsumer("order", "ch01", config)
	if err != nil {
		log.Fatalln(err)
	}
	// 添加消息处理器
	consumer.AddHandler(&MyHandler{})

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
