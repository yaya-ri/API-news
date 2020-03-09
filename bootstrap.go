package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	controllers "github.com/yaya-ri/API-news/module/v1/controller"
	models "github.com/yaya-ri/API-news/module/v1/model"
	requests "github.com/yaya-ri/API-news/module/v1/object/request"

	service "github.com/yaya-ri/API-news/module/v1/service"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

//DB global variable
var DB *gorm.DB

var newsService service.NewsServiceInterface
var elasticSearchService service.ElasticSearchServiceInterface
var queueService service.QueueServiceInterface

//RouteInit initial route apps
func RouteInit() {

	DB.AutoMigrate(models.News{})

	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPort, _ := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	rabbitMQUser := os.Getenv("RABBITMQ_USER")
	rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitMQQueue := "news"
	esHost := os.Getenv("ES_HOST")
	esPort, _ := strconv.Atoi(os.Getenv("ES_PORT"))

	newsController := controllers.NewsControllerHandler(DB, rabbitMQHost, rabbitMQPort, rabbitMQUser, rabbitMQPassword, rabbitMQQueue, esHost, esPort)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	store := persistence.NewInMemoryStore(time.Second)

	r.GET("/news", cache.CachePage(store, 5*time.Second, newsController.Find))
	r.POST("/news", newsController.Store)

	serverPort := os.Getenv("APP_PORT")
	log.Info("run on port ", serverPort)

	fmt.Printf("App running on port %s\n", serverPort)

	if err := http.ListenAndServe(":"+serverPort, r); err != nil {
		log.Fatal(err)
	}
}

//InitService initial for service
func InitService() {
	var err error

	esHost := os.Getenv("ES_HOST")
	esPort, _ := strconv.Atoi(os.Getenv("ES_PORT"))
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPort, _ := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	rabbitMQUser := os.Getenv("RABBITMQ_USER")
	rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitMQQueue := "news"

	newsService = service.NewsServiceHandler(DB)
	elasticSearchService = service.ElasticSearchServiceHandler(esHost, esPort)

	queueService, err = service.QueueServiceHandler(rabbitMQHost, rabbitMQPort, rabbitMQUser, rabbitMQPassword, rabbitMQQueue)
	if err != nil {
		log.Fatal(err)
	}

}

//InitElasticSearch initial for ES
func InitElasticSearch() {
	var err error

	index := "news"
	indexExists, err := elasticSearchService.CheckIndexExist(index)
	if err != nil {
		log.Fatal(err)
	}

	if !indexExists {
		err = elasticSearchService.CreateIndex(index, map[string]interface{}{
			"mappings": map[string]interface{}{
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "integer",
					},
					"created": map[string]interface{}{
						"type":   "date",
						"format": "yyyy-MM-dd HH:mm:ss",
					},
				},
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

//QueueListener godoc
func QueueListener() {
	msgs, err := queueService.Subcribe()
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			payload := requests.News{}
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Fatal(err)
			}

			_, err = StoreNews(payload)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Success store news!!!")
		}
	}()

	fmt.Println("Waiting for messages")
	<-forever
}

//StoreNews store news to mysql and elasticsearch
func StoreNews(request requests.News) (models.News, error) {
	result, err := newsService.Store(request)
	if err != nil {
		return models.News{}, err
	}

	err = elasticSearchService.Store("news", fmt.Sprintf("%d", result.ID), map[string]interface{}{
		"id":      result.ID,
		"created": result.Created.Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return models.News{}, err
	}
	return result, nil
}
