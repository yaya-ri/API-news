package services

import (
	"fmt"
	"testing"
	"time"
)

func TestCheckIndexExist(t *testing.T) {
	service := ElasticSearchServiceHandler("localhost", 9201)
	service.Clear()

	expected := false

	res, err := service.CheckIndexExist("newss")
	if err != nil {
		t.Error("ElasticSearchService.IndexExists failed with error")
		t.Errorf(err.Error())
	}
	if res != expected {
		t.Error("ElasticSearchService.IndexExists failed to retrieve expected data")
	}
}

func TestCreateIndex(t *testing.T) {

	service := ElasticSearchServiceHandler("localhost", 9201)
	service.Clear()

	err := service.CreateIndex("newss", map[string]interface{}{
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
		t.Error("ElasticSearchService.AddIndex failed with error")
		t.Errorf(err.Error())
	}
}

func TestElasticSearchAddIndex(t *testing.T) {

	service := ElasticSearchServiceHandler("localhost", 9201)
	service.Clear()

	service.CreateIndex("newss", map[string]interface{}{
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

	payload := map[string]interface{}{
		"id":      1,
		"created": time.Now().Format("2006-01-02 15:04:05"),
	}

	err := service.Store("newss", fmt.Sprintf("%d", 1), payload)
	if err != nil {
		t.Error("ElasticSearchService.AddItem failed with error")
		t.Errorf(err.Error())
	}
}
