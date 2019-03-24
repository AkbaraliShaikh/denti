package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCtrl struct{}

func NewHealthCtrl() *healthCtrl {
	return &healthCtrl{}
}

func (h healthCtrl) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
