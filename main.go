package main

import (
	"github.com/yaya-ri/API-news/db"

	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
	os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	log.Info("loading .env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	log.Info("success loading .env")
	log.Info("Starting server... \n")

	env := os.Getenv("APP_ENV")
	if env == "development" {
		startApp()
	} else if env == "production" {
		startApp()
	} else {
		startApp()
	}

}

//startApp start API
func startApp() {
	dbConn, err := db.InitDBSQL()
	if err != nil {
		fmt.Println(err)
	}

	DB = dbConn

	InitService()
	InitElasticSearch()

	go QueueListener()
	RouteInit()

}
