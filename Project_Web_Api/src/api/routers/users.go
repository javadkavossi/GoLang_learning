package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/GoLang_learning/api/handlers"
	"github.com/javadkavossi/GoLang_learning/config"
)

func Users(router *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewUsersHandler(cfg)
	router.POST("/sendOtp", handler.SendOtp)
}
