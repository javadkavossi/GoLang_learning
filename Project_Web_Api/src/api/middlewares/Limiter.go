package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/GoLang_learning/api/helper"
)

func Limiter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithValidationError("To Many Requests !", false, http.StatusTooManyRequests, err))
			return
		} else {
			ctx.Next()
		}
	}
}
