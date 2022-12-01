package rest

import (
	"net/http"

	"github.com/akbaralishaikh/denti/pkg/logger"
	"github.com/akbaralishaikh/denti/pkg/login"

	"github.com/gin-gonic/gin"
)

type loginCtrl struct {
	log logger.LogInfoFormat
	svc login.Service
}

func NewLoginCtrl(log logger.LogInfoFormat, svc login.Service) *loginCtrl {
	return &loginCtrl{log, svc}
}

func (l loginCtrl) Signin(ctx *gin.Context) {
	var lg login.Login
	if err := ctx.ShouldBindJSON(&lg); err != nil {
		l.log.Errorf("signin request bind failed with reason : %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok := l.svc.Signin(lg.Email, lg.Password); !ok {
		ctx.Status(http.StatusForbidden)
		return
	}

	ctx.Status(http.StatusOK)
}
