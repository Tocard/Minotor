package utils

import (
	"2miner-monitoring/config"
	"github.com/gofrs/flock"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
)

func HandleHttpError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func HandleFatalMsg(msg string) {
	log.Fatalln(msg)
}

func join(s ...string) string {
	return strings.Join(s, "")
}

func IsValidAdresse(wallet string) bool {
	//TODO: do something smart & pertinent here
	return true
}

func WriteYaml() bool {
	data, err := yaml.Marshal(config.Cfg.Adress)
	if err != nil {
		log.Fatal(err)
	}
	fileLock := flock.New(config.Cfg.LockPath)
	locked, err := fileLock.TryLock()
	if err != nil {
		return false
	}
	if locked {
		err = ioutil.WriteFile(config.Cfg.AdressFilePath, data, 0)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Adresse updated")
	}
	fileLock.Unlock()
	return true
}
