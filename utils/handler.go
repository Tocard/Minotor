package utils

import (
	"2miner-monitoring/config"
	"bytes"
	"encoding/json"
	"github.com/gofrs/flock"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func HandleHttpError(err error) {
	if err != nil {
		log.Printf("Error is %s", err)
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
	data, err := yaml.Marshal(config.Wtw)
	if err != nil {
		log.Printf("Not able to Marshal Adress Config File")
		return false
	}
	fileLock := flock.New(config.Cfg.LockPath)
	locked, err := fileLock.TryLock()
	if err != nil {
		log.Printf("Not able to Lock Adress file")
		return false
	}
	if locked {
		err = ioutil.WriteFile(config.Cfg.AdressFilePath, data, 0)
		if err != nil {
			log.Printf("Not able to Write Adress to file")
			return false
		}
		log.Printf("Adresse updated")
	}
	err = fileLock.Unlock()
	if err != nil {
		log.Printf("Not able to UnLock Adress file")
		return false
	}
	return true
}

// Execute a request. If token provided, it's added as Authorization.
func DoRequest(method, url string, data interface{}, token ...string) (*http.Response, error) {
	b, _ := json.Marshal(data)
	body := bytes.NewReader(b)
	req, _ := http.NewRequest(method, url, body)
	log.Printf("body Dorequest %s", body)

	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token[0])
	}
	log.Printf("Req = %s", req)
	log.Printf("doqurest token = %s", token[0])
	client := http.Client{}
	return client.Do(req)
}
