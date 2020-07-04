package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
)

var logger = logrus.New()

func main() {

	fmt.Println("Test logging")

	logFile, err := os.OpenFile("cl-deploy.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()
	logger.SetOutput(logFile)
	logger.Error("Some error, continue")
	logger.Warn("Warning....crash to ground imminent.")
	logger.Debug("Just debugging information.")
	logger.Level = logrus.DebugLevel
	logger.Debug("Just debugging information.")
	logger.WithFields(logrus.Fields{"pkgdir": "/usr/local/18.31", "config": "/usr/18.31/config/deploy-config.json"}).Debug("debug info")
	logger.WithFields(logrus.Fields{"pkgdir": "/usr/local/18.31", "config": "/usr/18.31/config/deploy-config.json"}).Info("debug info")
	logger.Info("Simple, Info")
	logger.Fatal("Abort!")
	logger.Panic("Panic!")
}
