package repositories

import (
	models "github.com/yaya-ri/API-news/module/v1/model"
)

//NewsRepositoryInterface contract for news repository
type NewsRepositoryInterface interface {
	Store(model *models.News) error
	Find(id uint) (models.News, error)
}
