package data

type Hopers struct {
	Timestamp string `json:"@timestamp"`
	Data      struct {
		Balance int `json:"balance"`
	} `json:"data"`
}
