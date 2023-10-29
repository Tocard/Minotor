package ChiaDbPoolData

type Model struct {
	EsTimestamp string `json:"@timestamp"`
	FetchedAt   string `json:"fetched_at"`
	FetchedBy   string `json:"fetched_by"`
}

type Tabler interface {
	TableName() string
}
