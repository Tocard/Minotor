package utils

import (
	"bytes"
	"encoding/json"
	"log"
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

// Execute a request. If token provided, it's added as Authorization.
func DoRequest(method, url string, data interface{}, token ...string) (*http.Response, error) {
	b, _ := json.Marshal(data)
	body := bytes.NewReader(b)
	req, _ := http.NewRequest(method, url, body)
	if token != nil && len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token[0])
		req.BasicAuth()
	}
	client := http.Client{}
	return client.Do(req)
}
