package pullmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func PullMQ() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		panic("1111" + err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		panic("2222" + err.Error())
	}

	// err = ch.ExchangeDeclare(
	// 	"logs_direct",
	// 	"direct",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// if err != nil {
	// 	panic("3333" + err.Error())
	// }

	que, err := ch.QueueDeclare(
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic("3333" + err.Error())
	}

	// err = ch.QueueBind(
	// 	que.Name,
	// 	"orange",
	// 	"logs_direct",
	// 	false,
	// 	nil,
	// )
	// if err != nil {
	// 	panic("123" + err.Error())
	// }

	msgs, err := ch.Consume(
		que.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic("4444" + err.Error())
	}

	var fo chan struct{}

	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body), "----")
		}
	}()
	fmt.Println("结束了")
	<-fo
}

