package main

import (
	"log"

	"github.com/javadkavossi/GoLang_learning/api"
	"github.com/javadkavossi/GoLang_learning/config"
	"github.com/javadkavossi/GoLang_learning/data/cache"
	"github.com/javadkavossi/GoLang_learning/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()

	if err != nil {
		log.Fatal("Error Redis : ", err)
	}

	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal("Error Postgres : ", err)
	}
	defer db.CloseDb()

	api.InitServer(cfg)

}

// Add Swager ...
//go install github.com/swaggo/swag/cmd/swag@latest
// go get github.com/swaggo/gin-swagger
//go get github.com/swaggo/swag
// go get github.com/swaggo/files
