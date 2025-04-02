package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Cfg *Config

type Config struct {
	ElasticsearchUser     string   `yaml:"elasticsearch_user"`
	ElasticsearchPassword string   `yaml:"elasticsearch_password"`
	ElasticsearchHosts    []string `yaml:"elasticsearch_hosts"`
	LogLevel              string   `yaml:"log_level"`
	CaPath                string   `yaml:"ca_path"`

	APIAdress        string `yaml:"api_adress"`
	APIPort          int    `yaml:"api_port"`
	AutonomysNodeUrl string `yaml:"autonomys_node_url"`
	APItemplatesPath string `yaml:"api_templates_path"`
}

func LoadYamlConfig(ConfigFilePath string) {
	t := Config{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalln(err)
	}
	Cfg = &t
}
