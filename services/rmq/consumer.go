package rmq

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"order-api/utils"
	"time"
)

type Consumer struct {
	Connection  *amqp.Connection
	Channel     *amqp.Channel
	Queue       amqp.Queue
	MessageChan chan []byte
}

func NewConsumer(connection *amqp.Connection, exchange, routingKey, queueName string, ch chan []byte) *Consumer {
	channel, err := connection.Channel()
	utils.IsError(err, "failed open channel")
	err = channel.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.IsError(err, "failed exchange declare")
	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.IsError(err, "failed queue declare")
	err = channel.QueueBind(
		queue.Name,
		routingKey,
		exchange,
		false,
		nil,
	)
	utils.IsError(err, "failed queue bind")
	return &Consumer{
		Connection:  connection,
		Channel:     channel,
		Queue:       queue,
		MessageChan: ch,
	}
}

func (c *Consumer) ConsumeMessages() {
	mes, err := c.Channel.Consume(
		c.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.IsError(err, "Failed to register a consumer")

	go func() {
		log.Printf(c.Queue.Name)
		for d := range mes {
			c.MessageChan <- d.Body
		}
	}()
}

func (c *Consumer) GetMessageByCondition(condition func(message []byte) bool, seconds time.Duration) []byte {
	for {
		select {
		case message := <-c.MessageChan:
			// Обработка полученного сообщения
			if condition(message) {
				return message
			}
		case <-time.After(seconds * time.Second):
			log.Error("response wasn't received")
			return nil
		}
	}

	return nil
}

/*
func (c *Consumer) GetMessageByCondition(condition func(message []byte) bool, test chan []byte, seconds time.Duration) []byte {
	for {
		select {
		case message := <-c.MessageChan:
			if condition(message) {
				select {
				case test <- message:
					return <-test
				case <-time.After(seconds * time.Second):
					log.Error("response wasn't received")
					return nil
				}
			}
		case <-time.After(seconds * time.Second):
			log.Error("condition wasn't met")
			return nil
		}
	}
}*/

func (c *Consumer) Close() {
	err := c.Channel.Close()
	utils.IsError(err, "failed close channel")

	err = c.Connection.Close()
	utils.IsError(err, "failed close connection")
}
