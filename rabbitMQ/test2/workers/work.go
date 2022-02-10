package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("work连接RabbitMQ失败:%s\n", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("建立信道失败:%s\n", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"taskQueue", // name
		true,        // 声明为持久队列
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		fmt.Printf("声明队列失败:%s\n", err)
		return
	}

	//建立一个Delivery的通道
	msgs, err := ch.Consume(
		q.Name, //队列名称
		"",     //消费者
		false,  //自动确认消息传递
		false,  //专有性
		false,  //是否保存本地
		false,  //是否等待
		nil,    //args
	)
	if err != nil {
		fmt.Printf("消费者信道建立失败%s\n", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("收到消息: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("完成\n")
			d.Ack(false) //手动确认消息传递
		}
	}()
	log.Printf("等待消息中,CTRL+C退出")
	<-forever
}
