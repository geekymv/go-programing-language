package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.Config{}
	producer, err := nsq.NewProducer("127.0.0.1:4150", &config)
	if err != nil {
		log.Fatalln(err)
	}
	defer producer.Stop()

	err = producer.Publish("topic-demo-01", []byte("hello-01"))
	if err != nil {
		log.Fatalln(err)
	}
}
