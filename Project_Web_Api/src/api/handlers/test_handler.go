package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/GoLang_learning/api/helper"
)

type header struct {
	UserId      string
	Browser     string
	ContentType string
	Host        string
}

type personData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=3,max=50"`
	LastName     string `json:"last_name" binding:"required,alpha,min=3,max=50"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile"`
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Working ! ", true, http.StatusOK))
}

func (h *TestHandler) Users(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("user", true, http.StatusOK))
}

func (h *TestHandler) UserById(ctx *gin.Context) {
	id := ctx.Param("id")
	result := map[string]string{
		"result": "User By Id",
		"id":     id,
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(result, true, http.StatusOK))
}

func (h *TestHandler) UserName(ctx *gin.Context) {
	userName := ctx.Param("username")
	ctx.JSON(http.StatusOK, gin.H{
		"result":   "User Name",
		"UserName": userName,
	})
}

func (h *TestHandler) Accounts(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result":   "Account",
		"UserName": id,
	})
}

func (h *TestHandler) AddUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Add User",
	})
}

func (h *TestHandler) HeaderBinder1(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	contentType := ctx.GetHeader("Content-Type")

	ctx.JSON(http.StatusOK, gin.H{
		"result":      "HeaderBinder1",
		"UserId":      userId,
		"contentType": contentType,
	})
}

func (h *TestHandler) HeaderBinder2(ctx *gin.Context) {
	header := header{}
	ctx.BindHeader(&header)
	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"Header": header,
	})
}

func (h *TestHandler) QueryBinder1(ctx *gin.Context) {
	name := ctx.Query("name")
	family := ctx.Query("family")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder 1",
		"name":   name,
		"family": family,
	})
}

func (h *TestHandler) QueryBinder2(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	name := ctx.Query("name")
	family := ctx.Query("family")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder 2",
		"name":   name,
		"family": family,
		"ids":    ids,
	})
}

func (h *TestHandler) UriBinder(ctx *gin.Context) {
	data := ctx.Param("Data")
	ctx.JSON(http.StatusOK, gin.H{
		"result":   "User Name",
		"UserName": data,
	})
}

func (h *TestHandler) BodyBinder(ctx *gin.Context) {
	person := personData{}
	err := ctx.ShouldBindJSON(&person)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError("Bad Request !", false, http.StatusBadRequest, err))
		return
	}
	// result := map[any]any{
	// 	"result":     "BodyBinder",
	// 	"personData": person,
	// }

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(person, true, http.StatusOK))
}

func (h *TestHandler) FormBinder(ctx *gin.Context) {
	person := personData{}
	err := ctx.ShouldBind(&person)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError("Bad Request !", false, http.StatusBadRequest, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(person, true, http.StatusOK))
}

func (h *TestHandler) FileBinder(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	err := ctx.SaveUploadedFile(file, "file.png")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result":   "BodyBinder",
		"FileName": file.Filename,
	})
}
