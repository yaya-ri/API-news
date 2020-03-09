package repositories

import (
	models "github.com/yaya-ri/API-news/module/v1/model"

	"github.com/jinzhu/gorm"
)

//NewsRepository godoc
type NewsRepository struct {
	DB        *gorm.DB
	TableName string
}

//NewsRepositoryHandler godoc
func NewsRepositoryHandler(db *gorm.DB) NewsRepositoryInterface {
	return &NewsRepository{
		DB:        db,
		TableName: "news",
	}
}

//Store save data news into table news
func (repo *NewsRepository) Store(model *models.News) error {
	return repo.DB.Save(model).Error
}

//Find search data from table news
func (repo *NewsRepository) Find(id uint) (models.News, error) {
	response := models.News{}
	query := repo.DB.Where("id = ?", id).Find(&response)
	return response, query.Error
}
