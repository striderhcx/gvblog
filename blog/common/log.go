package common

import (
	"os"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()


func init() {
	file, err := os.OpenFile("web.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}

	// Todo: log增加trace_id,跟踪整个请求调用链
	//Log.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")
}
