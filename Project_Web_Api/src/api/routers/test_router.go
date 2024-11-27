package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/GoLang_learning/api/handlers"
)


func TestRouter(router *gin.RouterGroup) {
	h := handlers.NewTestHandler()
	router.GET("/", h.Test)
	router.GET("/users", h.Users)
	router.GET("/user/:id", h.UserById)
	router.GET("/user/get-user-by-username/:username", h.UserName)
	router.GET("/user/:id/accounts", h.Accounts)
	router.POST("/HeaderBinder1", h.HeaderBinder1)
	router.POST("/HeaderBinder2", h.HeaderBinder2)
	router.POST("/QueryBinder1", h.QueryBinder1)
	router.POST("/QueryBinder2", h.QueryBinder2)
	router.POST("/UriBinder/:Data", h.UriBinder)
	router.POST("/BodyBinder", h.BodyBinder)
	router.POST("/FormBinder", h.FormBinder)
	router.POST("/file-binder", h.FileBinder)

}
