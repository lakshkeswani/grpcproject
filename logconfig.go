package grpcproject

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"

	"os"
	"runtime"
	"strings"
)

type LogX struct {
}

func (receiver LogX) Debug(agrs ...interface{}) {
	logrus.Debug(agrs...)
}

func (receiver LogX) Trace(agrs ...interface{}) {
	logrus.Info(agrs...)
}

func (receiver LogX) Warn(agrs ...interface{}) {
	logrus.Warn(agrs...)
}

func (receiver LogX) Error(agrs ...interface{}) {
	logrus.Error(agrs...)
}

func (receiver LogX) Info(agrs ...interface{}) {
	logrus.Info(agrs...)
}

func (receiver LogX) Fatal(agrs ...interface{}) {
	logrus.Info(agrs...)
}

func (receiver LogX) Panic(input string) {
	logrus.Panic(input)
}
func (receiver LogX) Initalizer() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmet everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logrus.SetOutput(file)
	//level of log file
	//	log.SetLevel(log.DebugLevel)
}
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

var Log LogX
