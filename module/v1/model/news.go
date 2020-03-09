package models

import "time"

//News reference news
type News struct {
	//gorm.Model
	ID      uint
	Author  string
	Body    string
	Created time.Time
}
