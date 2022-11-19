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
	APITokenEtherscan     string   `yaml:"api_token_etherscan"`
	LogLevel              string   `yaml:"log_level"`
	CaPath                string   `yaml:"ca_path"`
	RedisHost             string   `yaml:"redis_host"`
	RedisPassword         string   `yaml:"redis_password"`
	MinerListing          string   `yaml:"miner_listing"`
	APILogFile            string   `yaml:"api_log_file"`
	APIUsername           string   `yaml:"api_username"`
	APIPassword           string   `yaml:"api_password"`
	APIAdress             string   `yaml:"api_adress"`
	AdressFilePath        string   `yaml:"adress_file_path"`
	LockPath              string   `yaml:"lock_path"`
	CoinList              []string `yaml:"coin_list"`
	ElasticsearchHosts    []string `yaml:"elasticsearch_hosts"`
	ElasticsearchPort     int      `yaml:"elasticsearch_port"`
	RedisPort             int      `yaml:"redis_port"`
	RedisShortLifetime    int      `yaml:"redis_short_lifetime"`
	RedisMidLifetime      int      `yaml:"redis_mid_lifetime"`
	RedisLongLifetime     int      `yaml:"redis_long_lifetime"`
	APIPort               int      `yaml:"api_port"`
	APIFrontPort          int      `yaml:"api_front_port"`
	Factor                float64  `yaml:"factor"`
	EthFactor             float64  `yaml:"ether_factor"`
	GazFactor             float64  `yaml:"gaz_factor"`
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
