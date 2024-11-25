package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/javadkavossi/GoLang_learning/api/middlewares"
	"github.com/javadkavossi/GoLang_learning/api/routers"
	"github.com/javadkavossi/GoLang_learning/api/validations"
	"github.com/javadkavossi/GoLang_learning/config"
)

func InitServer() {

	cfg := config.GetConfig()

	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}

	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.Limiter())

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.TestRouter(test_router)
		routers.Health(health)
	}

	v2 := api.Group("/v2")
	{
		test_router2 := v2.Group("/test")
		routers.TestRouter(test_router2)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}
