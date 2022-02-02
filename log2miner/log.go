package log2miner

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLogger(LogFilePath string) {
	stdout := os.Stdout
	if len(LogFilePath) > 0 {
		var err error
		log.Println("Current log2miner is gonna be written in " + LogFilePath)
		stdout, err = os.OpenFile(LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	InfoLogger = log.New(stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func ErrorMessage(msg error) {
	if msg != nil {
		ErrorLogger.Println(msg.Error())
	}
}

func WarnErrorMessage(msg error) {
	if msg != nil {
		WarningLogger.Println(msg.Error())
	}
}

func Error(msg interface{}) {
	if msg != nil {
		ErrorLogger.Println(msg)
	}
}

func Warn(msg interface{}) {
	if msg != nil {
		WarningLogger.Println(msg)
	}
}

func Log(msg interface{}) {
	if msg != nil {
		InfoLogger.Println(msg)
	}
}
