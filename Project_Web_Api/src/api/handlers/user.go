package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/GoLang_learning/api/dto"
	"github.com/javadkavossi/GoLang_learning/api/helper"
	"github.com/javadkavossi/GoLang_learning/config"
	"github.com/javadkavossi/GoLang_learning/services"
)

type UserHandler struct {
	services *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UserHandler {
	return &UserHandler{
		services: services.NewUserService(cfg),
	}
}

// Send OTP
// @Summary Send OTP
// @Description Send an OTP to a user
// @Tags User
// @Accept json
// @produces json
// @Param Request body dto.GetOtpRequest true "Get Otp Request"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/users/otp [Post]
func (h *UserHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err))
		return
	}
	err = h.services.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusInternalServerError, err))
		return
	}

	//Call internal SMS Service

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(map[string]string{
		"result": "Otp Sent",
	}, true, http.StatusOK))
}
