package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/yaya-ri/API-news/module/v1/model"
)

//ElasticSearchService godoc
type ElasticSearchService struct {
	host string
	port int
	http *http.Client
}

//ElasticSearchServiceHandler godoc
func ElasticSearchServiceHandler(host string, port int) ElasticSearchServiceInterface {
	return &ElasticSearchService{
		host: host,
		port: port,
		http: &http.Client{},
	}
}

//CheckIndexExist check if index exist in elasticsearch
func (service *ElasticSearchService) CheckIndexExist(name string) (bool, error) {
	url := fmt.Sprintf("http://%s:%d/%s", service.host, service.port, name)

	request, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return false, err
	}

	res, err := service.http.Do(request)
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, err
	}

	return true, err
}

//CreateIndex create new index for elasticsearch
func (service *ElasticSearchService) CreateIndex(name string, maps map[string]interface{}) error {
	url := fmt.Sprintf("http://%s:%d/%s", service.host, service.port, name)
	bodyRequest, err := json.Marshal(maps)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyRequest))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	res, err := service.http.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed create index %s", name)
	}
	return nil
}

//Store save document to elasticsearch
func (service *ElasticSearchService) Store(index, ID string, doc map[string]interface{}) error {
	url := fmt.Sprintf("http://%s:%d/%s/_doc/%s", service.host, service.port, index, ID)
	bodyRequest, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyRequest))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	res, err := service.http.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("Failed save document elasticsearch %s", index)
	}
	return nil
}

//Find get list data from elasticsearch
func (service *ElasticSearchService) Find(index string, size, from int, filter map[string]interface{}) ([]models.ElasticSearch, error) {
	url := fmt.Sprintf("http://%s:%d/%s/_search?size=%d&from=%d", service.host, service.port, index, size, from)
	bodyRequest, err := json.Marshal(filter)
	if err != nil {
		return []models.ElasticSearch{}, err
	}

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(bodyRequest))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return []models.ElasticSearch{}, err
	}

	res, err := service.http.Do(req)
	if err != nil {
		return []models.ElasticSearch{}, err
	}

	var result models.ElasticSearchResult
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return []models.ElasticSearch{}, err
	}

	return result.Result.List, nil
}

// Clear delete all index
func (service *ElasticSearchService) Clear() error {
	url := fmt.Sprintf("http://%s:%d/_all", service.host, service.port)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = service.http.Do(req)
	if err != nil {
		return err
	}

	return nil
}
