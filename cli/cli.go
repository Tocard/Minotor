package cli

import (
	"flag"
)

type CliStruct struct {
	FilePathConfig string
}

func Cli() CliStruct {
	config := CliStruct{}
	// Defining option flags. For this, we're using the Flag package from the standard library
	// We need to define three arguments: the flag's name, the default value, and a short description (displayed whith the option --help)
	flag.StringVar(&config.FilePathConfig, "config", "config/config.yml", "file config path")
	flag.Parse() // This will parse all the arguments from the terminal
	return config
}
