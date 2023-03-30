package cli

import (
	"flag"
)

type Struct struct {
	FilePathConfig string
}

func Cli() Struct {
	config := Struct{}
	flag.StringVar(&config.FilePathConfig, "config", "/etc/minotor/config.yml", "file config path")
	flag.Parse()
	return config
}
