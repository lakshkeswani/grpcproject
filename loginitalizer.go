package grpcproject

import (
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func LogInitalizer() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type Customlog struct {
}

func (receiver Customlog) Infolog(input string) {
	infoLogger.Println(input)
}
func (receiver Customlog) Warninglog(input string) {
	warningLogger.Println(input)
}
func (receiver Customlog) Errorlog(input string) {
	errorLogger.Println(input)
}
