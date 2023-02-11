package data

type Hopers struct {
	Timestamp string `json:"@timestamp"`
	Data      struct {
		Balance string `json:"balance"`
	} `json:"data"`
}
