package main

import (
	"github.com/streadway/amqp"
	"log"
)

//检查amqp调用的返回值
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//连接RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "连接RabbitMQ失败")
	defer conn.Close()

	//声明一个RabbitMQ信道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//声明一个RabbitMQ队列,并配置参数
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   //durable
		false,   //delete when unused
		false,   // exclusive
		false,   //no-wait
		nil,     //argument
	)
	failOnError(err, "声明RabbitMQ队列失败")

	body := "Hello World!"
	// 4.将消息发布到声明的队列
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
