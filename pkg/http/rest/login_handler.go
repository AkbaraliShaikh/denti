package rest

import (
	"net/http"

	"denti/pkg/logger"
	"denti/pkg/login"

	"github.com/gin-gonic/gin"
)

type loginController struct {
	log logger.LogInfoFormat
	svc login.Service
}

func NewLoginController(log logger.LogInfoFormat, svc login.Service) *loginController {
	return &loginController{log, svc}
}

func (l loginController) Signin(ctx *gin.Context) {
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
