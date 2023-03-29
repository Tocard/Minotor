package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"minotor/config"
	"net/http"
	"strings"
)

func HandleHttpError(err error) {
	if err != nil {
		log.Printf("Error is %s", err)
	}
}

func join(s ...string) string {
	return strings.Join(s, "")
}

func IsValidAdresse(wallet string) bool {
	//TODO: do something smart & pertinent here
	return true
}

// Execute a request. If token provided, it's added as Authorization.
func DoRequest(method, url string, data interface{}, token ...string) (*http.Response, error) {
	b, _ := json.Marshal(data)
	body := bytes.NewReader(b)
	req, _ := http.NewRequest(method, url, body)
	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token[0])
	}
	req.BasicAuth()
	client := http.Client{}
	return client.Do(req)
}

func DoRequestAuth(method, url string, data interface{}) (*http.Response, error) {
	b, _ := json.Marshal(data)
	body := bytes.NewReader(b)
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(config.Cfg.GrafanaUser, config.Cfg.GrafanaPassword)
	client := http.Client{}
	return client.Do(req)
}
