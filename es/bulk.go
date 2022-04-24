package es

import (
	"context"
	"github.com/dustin/go-humanize"
	"log"
	"strings"
	"time"
)

func BulkData(indexName string, DataToShip [][]byte) {
	var (
		countSuccessful uint64
	)
	start := time.Now().UTC()
	for _, JsonRaw := range DataToShip {
		AddIemIndexer(Indexers[indexName], JsonRaw, &countSuccessful)
	}
	if err := Indexers[indexName].Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}
	biStats := Indexers[indexName].Stats()
	log.Println(strings.Repeat("▔", 65))
	dur := time.Since(start)
	if biStats.NumFailed > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	}
}
