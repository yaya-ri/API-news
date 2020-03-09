package requests

//News requests json format
type News struct {
	Author string `json:"author" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

//NewsResponse godoc
type NewsResponse struct {
	ID      uint   `json:"id"`
	Author  string `json:"author"`
	Body    string `json:"body"`
	Created string `json:"created"`
}
