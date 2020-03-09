package controllers

import (
	"net/http"
	"sort"
	"sync"

	helpers "github.com/yaya-ri/API-news/module/v1/helper"
	models "github.com/yaya-ri/API-news/module/v1/model"
	requests "github.com/yaya-ri/API-news/module/v1/object/request"
	service "github.com/yaya-ri/API-news/module/v1/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//NewsController godoc
type NewsController struct {
	newsService          service.NewsServiceInterface
	queueService         service.QueueServiceInterface
	elasticSearchService service.ElasticSearchServiceInterface
	helpers.ResponseHelper
}

//NewsControllerHandler godoc
func NewsControllerHandler(db *gorm.DB, RBMQhost string, RBMQport int, RBMQuser string, RBMQpassword string, RBMQqueue string, EShost string, ESport int) NewsControllerInterface {
	queueHanler, _ := service.QueueServiceHandler(RBMQhost, RBMQport, RBMQuser, RBMQpassword, RBMQqueue)
	return &NewsController{
		newsService:          service.NewsServiceHandler(db),
		queueService:         queueHanler,
		elasticSearchService: service.ElasticSearchServiceHandler(EShost, ESport),
	}
}

//Store store controller
func (controller *NewsController) Store(c *gin.Context) {
	var request requests.News
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response := controller.ResponseHelper.StatusBadRequest(gin.H{}, err.Error())
		c.JSON(response.Code, response)
		return
	}

	err = controller.queueService.Publish(request)
	if err != nil {
		response := controller.ResponseHelper.InternalServerError(gin.H{}, err.Error())
		c.JSON(response.Code, response)
		return
	}

	response := controller.ResponseHelper.SuccessResponse(gin.H{}, "Success store news")
	c.JSON(http.StatusOK, response)

}

//Find find controller
func (controller *NewsController) Find(c *gin.Context) {
	filter := map[string]interface{}{
		"sort": []map[string]interface{}{
			map[string]interface{}{
				"created": map[string]interface{}{
					"order": "desc",
				},
			},
		},
	}

	getES, err := controller.elasticSearchService.Find("news", 10, 0, filter)
	if err != nil {
		response := controller.ResponseHelper.InternalServerError(gin.H{}, err.Error())
		c.JSON(response.Code, response)
		return
	}

	newsChannel := make(chan models.News, len(getES))
	var wg sync.WaitGroup

	for _, data := range getES {
		wg.Add(1)
		newsID := uint(data.Source["id"].(float64))
		go func(ID uint, channel chan models.News, waitGroup *sync.WaitGroup) {
			news, err := controller.newsService.Find(ID)
			if err != nil {
				response := controller.ResponseHelper.InternalServerError(gin.H{}, err.Error())
				c.JSON(response.Code, response)
				return
			}

			defer waitGroup.Done()
			channel <- news
		}(newsID, newsChannel, &wg)
	}

	wg.Wait()
	close(newsChannel)

	newsList := make([]models.News, 0)
	for item := range newsChannel {
		newsList = append(newsList, item)
	}

	sort.Slice(newsList, func(i, j int) bool {
		return newsList[j].Created.Before(newsList[i].Created)
	})

	newsListResponse := make([]requests.NewsResponse, 0)
	for _, item := range newsList {
		newsListResponse = append(newsListResponse, requests.NewsResponse{
			ID:      item.ID,
			Author:  item.Author,
			Body:    item.Body,
			Created: item.Created.Format("2006-01-02 15:04:05"),
		})
	}

	response := controller.ResponseHelper.SuccessResponse(&newsListResponse, "Success get news")
	c.JSON(response.Code, response)
	return
}
