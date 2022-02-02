package utils

import "log"

func HandleHttpError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func HandleFatalMsg(msg string) {
	log.Fatalln(msg)
}
