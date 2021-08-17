package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rallinator7/akita-poker/server/checker"
	"github.com/rallinator7/akita-poker/server/controller"
	"github.com/rallinator7/akita-poker/server/logger"
	"go.uber.org/zap"
)

func main() {
	ev := map[string]string{
		"PORT": "",
	}
	for k := range ev {
		ev[k] = os.Getenv(k)
		if ev[k] == "" {
			log.Fatalf("environment variable %s is not set. Must be set in order to run", k)
		}
	}

	loggerMgr := logger.InitLogger()
	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync()

	r := gin.Default()

	handChecker := checker.NewHandChecker()
	handController := controller.NewHandController(handChecker)

	r.POST("/check", handController.PostHandCheck)

	r.Run(fmt.Sprintf(":%s", ev["PORT"]))
}
