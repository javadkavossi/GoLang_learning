package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/GoLang_learning/api/handlers"
)

func Health(router *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	router.GET("/", handler.Health)
	router.GET("/:id", handler.HealthGwtById)
}
