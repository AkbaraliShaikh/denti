package main

import (
	"fmt"
	"os"

	"github.com/akbaralishaikh/denti/cmd/server"
	"github.com/akbaralishaikh/denti/pkg/di"
	"github.com/akbaralishaikh/denti/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {

	g := gin.Default()
	d := di.BuildContainer()

	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})

	svr := server.NewServer(g, d, l)
	svr.MapRoutes()
	if err := svr.SetupDB(); err != nil {
		return err
	}
	return svr.Start()
}
