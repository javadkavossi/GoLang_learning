package handlers

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
func (h *HealthHandler) Health(ctx *gin.Context) {
	var sum float64
	for i := int64(0); i < 1e9; i++ {
		sum += math.Sqrt(float64(i))
		sum += sum
	}
	ctx.JSON(http.StatusOK, "Working")
	return

}

func (h *HealthHandler) HealthGwtById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, fmt.Sprintf("Working By Id : %s", id))

}
