package services

import (
	"time"

	models "github.com/yaya-ri/API-news/module/v1/model"
	requests "github.com/yaya-ri/API-news/module/v1/object/request"
	repository "github.com/yaya-ri/API-news/module/v1/repository"

	"github.com/jinzhu/gorm"
)

//NewsService godoc
type NewsService struct {
	DB   *gorm.DB
	News repository.NewsRepositoryInterface
}

//NewsServiceHandler godoc
func NewsServiceHandler(db *gorm.DB) NewsServiceInterface {
	return &NewsService{
		DB:   db,
		News: repository.NewsRepositoryHandler(db),
	}
}

//Store service for call repository store
func (service *NewsService) Store(request requests.News) (models.News, error) {
	model := models.News{
		Author:  request.Author,
		Body:    request.Body,
		Created: time.Now(),
	}

	return model, service.News.Store(&model)
}

//Find service for call repository find
func (service *NewsService) Find(id uint) (models.News, error) {
	response, err := service.News.Find(id)
	return response, err
}
