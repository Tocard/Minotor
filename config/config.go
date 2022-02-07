package config

import (
	"2miner-monitoring/log2miner"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Cfg *Config

type Config struct {
	ElasticsearchUser     string   `yaml:"elasticsearch_user"`
	ElasticsearchPassword string   `yaml:"elasticsearch_password"`
	ElasticsearchHosts    []string `yaml:"elasticsearch_hosts"`
	ElasticsearchPort     int      `yaml:"elasticsearch_port"`
	APITokenEtherscan     string   `yaml:"api_token_etherscan"`
	LogLevel              string   `yaml:"log_level"`
	CaPath                string   `yaml:"ca_path"`
	TwoMinersURL          string   `yaml:"2miners_url"`
	Adress                []string `yaml:"adress"`
	RedisHost             string   `yaml:"redis_host"`
	RedisPort             int      `yaml:"redis_port"`
	RedisPassword         string   `yaml:"redis_password"`
	RedisLifetime         int      `yaml:"redis_lifetime"`
	Factor                float64  `yaml:"factor"`
	EthFactor             float64  `yaml:"ether_factor"`
	GazFactor             float64  `yaml:"gaz_factor"`
	MinerListing          string   `yaml:"miner_listing"`
	APIPort               int      `yaml:"api_port"`
	APILogFile            string   `yaml:"api_log_file"`
	APIUsername           string   `yaml:"api_username"`
	APIPassword           string   `yaml:"api_password"`
}

func LoadYamlConfig(ConfigFilePath string) {
	t := Config{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	log2miner.Error(err)
	err = yaml.Unmarshal(data, &t)
	log2miner.Error(err)
	Cfg = &t
}
