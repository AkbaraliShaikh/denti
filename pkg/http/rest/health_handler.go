package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthController struct{}

func NewHealthController() *healthController {
	return &healthController{}
}

func (h healthController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
