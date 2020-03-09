package models

//ElasticSearch godoc
type ElasticSearch struct {
	Index  string                 `json:"_index"`
	Type   string                 `json:"_type"`
	ID     string                 `json:"_id"`
	Score  float32                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}

// ElasticSearchList godoc
type ElasticSearchList struct {
	List []ElasticSearch `json:"hits"`
}

// ElasticSearchResult godoc
type ElasticSearchResult struct {
	Result ElasticSearchList `json:"hits"`
}
