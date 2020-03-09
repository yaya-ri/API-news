package services

import (
	"github.com/yaya-ri/API-news/mocks"
	models "github.com/yaya-ri/API-news/module/v1/model"

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

// func TestNewsServiceStore(t *testing.T) {

// 	tests := []struct {
// 		name          string
// 		args          requests.News
// 		configureMock func(newsRepo *mocks.MockNewsRepositoryInterface)
// 		assertResult  func(response models.News, err error)
// 	}{
// 		{
// 			name: "Success create propose mitra",
// 			args: requests.News{
// 				Author: "yaya",
// 				Body:   "yaya",
// 			},
// 			configureMock: func(newsRepo *mocks.MockNewsRepositoryInterface) {
// 				model := models.News{
// 					Author:  "yaya",
// 					Body:    "yaya",
// 					Created: time.Now(),
// 				}

// 				newsRepo.EXPECT().Store(model).Return(nil).Times(1)
// 			},
// 			assertResult: func(response models.News, err error) {
// 				assert.Nil(t, err)
// 				//assert.Equal(t, 0, response)
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			db, _, err := sqlmock.New()
// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			database, _ := gorm.Open("mysql", db)
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			newsRepo := mocks.NewMockNewsRepositoryInterface(ctrl)
// 			tt.configureMock(newsRepo)
// 			service := NewsService{
// 				DB:   database,
// 				News: newsRepo,
// 			}

// 			model := models.News{
// 				Author:  "yaya",
// 				Body:    "yaya",
// 				Created: time.Now(),
// 			}

// 			_, err = service.Store(tt.args)
// 			tt.assertResult(model, err)
// 		})
// 	}

// }

func TestNewsServiceFind(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database, _ := gorm.Open("mysql", db)

	tests := []struct {
		name          string
		configureMock func(repository *mocks.MockNewsRepositoryInterface)
		assertResult  func(model models.News, err error)
	}{
		{
			name: "Success get news",
			configureMock: func(repository *mocks.MockNewsRepositoryInterface) {
				model := models.News{
					ID:      1,
					Author:  "yaya",
					Body:    "yaya",
					Created: time.Now(),
				}
				repository.EXPECT().Find(model.ID).Return(model, nil).Times(1)
			},
			assertResult: func(model models.News, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "yaya", model.Author)
				assert.Equal(t, "yaya", model.Body)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			newsRepo := mocks.NewMockNewsRepositoryInterface(ctrl)
			tt.configureMock(newsRepo)
			service := NewsService{
				DB:   database,
				News: newsRepo,
			}
			model, err := service.Find(1)
			tt.assertResult(model, err)
		})
	}
}
