package main

import (
	"log"

	"github.com/javadkavossi/GoLang_learning/api"
	"github.com/javadkavossi/GoLang_learning/config"
	"github.com/javadkavossi/GoLang_learning/data/cache"
	"github.com/javadkavossi/GoLang_learning/data/db"
	"github.com/javadkavossi/GoLang_learning/data/db/migrations"
	"github.com/javadkavossi/GoLang_learning/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
// @description Type "Bearer " followed by your token

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Starting Project", nil)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()

	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
		log.Fatal("Error Redis : ", err)
	}
	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
		log.Fatal("Error Postgres : ", err)
	}
	defer db.CloseDb()
	migrations.Up_1()
	api.InitServer(cfg)
}

// Add Swager ...
//go install github.com/swaggo/swag/cmd/swag@latest
// go get github.com/swaggo/gin-swagger
//go get github.com/swaggo/swag
// go get github.com/swaggo/files
