package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

// bodyFrom 从命令行获取将要发送的消息内容
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	}
	return s
}

func main() {
	//1.连接RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("连接RabbitMQ服务器失败:%s\n", err)
		return
	}
	defer conn.Close()

	//2.建立信道
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("建立信道失败:%s\n", err)
	}

	defer ch.Close()

	//3.声明要发送到的队列
	q, err := ch.QueueDeclare(
		"taskQueue", //要发送的队列名称
		true,        /*持久化,队列将一条消息发送给work以后会删除队列中的消息,持久化并不会删除队列中的消息可以等待work发送的确认消息
		,work如果没有响应则RabbitMQ会将此消息放回队列重新进行发送*/
		false, //关闭自动删除
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("声明队列失败:%s\n", err)
		return
	}

	body := bodyFrom(os.Args)
	//将消息发布到声明的队列
	err = ch.Publish(
		"",     //不选择交换器
		q.Name, //队列名称
		false,  //强制推送
		false,  //立即推送
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //持久化
			ContentType:  "text/plain",    //选择发布内容的类型
			Body:         []byte(body),    //选择发送内容将其装入一个byte数组中
		},
	)
	if err != nil {
		fmt.Printf("发布消息失败%s\n", err)
		return
	}
	log.Printf("消息已发送: %s", body)
}
