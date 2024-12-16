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
	"github.com/javadkavossi/GoLang_learning/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {

	r := gin.New()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	RegisterValidator()
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.Limiter())

	RegisterSwagger(r, cfg)

	RegisterRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

// ==================== Validator ======================

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}
}

// ==================== Router ======================

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
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
}

//  ==================== Swagger =======================

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang Web Api"
	docs.SwaggerInfo.Description = "golang Web Api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.UriHost, cfg.Server.ExternalPort)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
