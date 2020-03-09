package services

import (
	models "github.com/yaya-ri/API-news/module/v1/model"
	requests "github.com/yaya-ri/API-news/module/v1/object/request"

	"github.com/streadway/amqp"
)

//NewsServiceInterface contract for news repository
type NewsServiceInterface interface {
	Store(request requests.News) (models.News, error)
	Find(id uint) (models.News, error)
}

//ElasticSearchServiceInterface contract for elasaticsearch
type ElasticSearchServiceInterface interface {
	CheckIndexExist(name string) (bool, error)
	CreateIndex(name string, maps map[string]interface{}) error
	Store(index, ID string, doc map[string]interface{}) error
	Find(index string, size, from int, filter map[string]interface{}) ([]models.ElasticSearch, error)
	Clear() error
}

//QueueServiceInterface contract for RabbitMQ
type QueueServiceInterface interface {
	CreateChannel() (*amqp.Channel, error)
	CreateQueue(channel *amqp.Channel) (amqp.Queue, error)
	Publish(request interface{}) error
	Subcribe() (<-chan amqp.Delivery, error)
}
