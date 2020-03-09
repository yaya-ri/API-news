package services

import "testing"

func TestCheckIndexExist(t *testing.T) {
	service := ElasticSearchServiceHandler("localhost", 9207)
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
