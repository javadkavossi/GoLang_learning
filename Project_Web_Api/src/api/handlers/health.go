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

// HealthCheck godoc
// @Summary Check Health
// @Description Check the health of the server
// @Tags Health
// @Accept json
// @produces json
// @Param queryParam query string false "Example query parameter"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/health [Get]
func (h *HealthHandler) Health(ctx *gin.Context) {
	var sum float64

	for i := int64(0); i < 1e9; i++ {
		sum += math.Sqrt(float64(i))
		sum += sum
	}

	ctx.JSON(http.StatusOK, "Working")
	return

}

// HealthCheck godoc
// @Summary Check Health by Id
// @Description Check the health of the server By Id
// @Tags Health
// @Accept json
// @produces json
// @Param id path string false "ID of the health check"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/health/{id} [Get]
func (h *HealthHandler) HealthGwtById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, fmt.Sprintf("Working By Id : %s", id))

}
