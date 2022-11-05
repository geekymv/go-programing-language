package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {

	// Instantiate a producer.
	config := nsq.NewConfig()
	// 创建生产者
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		messageBody := []byte(fmt.Sprintf("hello-%d", i))
		err = producer.Publish("order", messageBody)
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
	}

	producer.Stop()

}
