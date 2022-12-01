package server

import (
	"fmt"

	"github.com/akbaralishaikh/denti/pkg/config"
	"github.com/akbaralishaikh/denti/pkg/logger"
	"github.com/akbaralishaikh/denti/pkg/patient"
	"github.com/akbaralishaikh/denti/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go.uber.org/dig"
)

type dserver struct {
	router *gin.Engine
	cont   *dig.Container
	logger logger.LogInfoFormat
}

func NewServer(e *gin.Engine, c *dig.Container, l logger.LogInfoFormat) *dserver {
	return &dserver{
		router: e,
		cont:   c,
		logger: l,
	}
}

func (ds *dserver) SetupDB() error {
	var db *gorm.DB

	if err := ds.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}

	db.Exec("SET search_path TO denti")
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&patient.Patient{})
	return nil
}

// Start start serving the application
func (ds *dserver) Start() error {
	var cfg *config.Config
	if err := ds.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return ds.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
