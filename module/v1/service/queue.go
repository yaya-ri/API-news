package services

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

//QueueService service for rabbitMQ
type QueueService struct {
	queue string
	amqp  *amqp.Connection
}

//QueueServiceHandler godoc
func QueueServiceHandler(host string, port int, user string, password string, queue string) (QueueServiceInterface, error) {
	amqpConn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		user,
		password,
		host,
		port),
	)

	if err != nil {
		return &QueueService{}, err
	}
	return &QueueService{
		queue: queue,
		amqp:  amqpConn,
	}, nil
}

//CreateChannel create channel for queue
func (service *QueueService) CreateChannel() (*amqp.Channel, error) {
	return service.amqp.Channel()
}

//CreateQueue godoc
func (service *QueueService) CreateQueue(channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		service.queue,
		true,
		false,
		false,
		false,
		nil,
	)
}

//Publish godoc
func (service *QueueService) Publish(request interface{}) error {
	chann, err := service.CreateChannel()
	if err != nil {
		return err
	}

	defer chann.Close()

	queue, err := service.CreateQueue(chann)
	if err != nil {
		return err
	}

	data, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = chann.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)

	return err
}

//Subcribe consume / subrice queue
func (service *QueueService) Subcribe() (<-chan amqp.Delivery, error) {
	chann, err := service.CreateChannel()
	if err != nil {
		return nil, err
	}

	queue, err := service.CreateQueue(chann)
	if err != nil {
		return nil, err
	}

	result, err := chann.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	return result, err
}
