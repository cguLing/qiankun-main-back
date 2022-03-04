package autoinit

import (
	"bus-backend-go/utils"
	"github.com/sirupsen/logrus"
	"os"
)

func InitTouch() {
}
func init() {
	initLog()
	//go cron.ServerSync()
}

func initLog() {
	utils.Log.Out = os.Stdout
	utils.Log.Formatter = &logrus.JSONFormatter{}
	utils.New(utils.Log)
	utils.Log.Info("bus-backend now start!")
}
