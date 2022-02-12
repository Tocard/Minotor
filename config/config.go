package config

import (
	"2miner-monitoring/log2miner"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Cfg *Config

type Adresses struct {
	Adress []string `yaml:"adress"`
}

type Config struct {
	ElasticsearchUser     string   `yaml:"elasticsearch_user"`
	ElasticsearchPassword string   `yaml:"elasticsearch_password"`
	ElasticsearchHosts    []string `yaml:"elasticsearch_hosts"`
	ElasticsearchPort     int      `yaml:"elasticsearch_port"`
	APITokenEtherscan     string   `yaml:"api_token_etherscan"`
	LogLevel              string   `yaml:"log_level"`
	CaPath                string   `yaml:"ca_path"`
	TwoMinersURL          string   `yaml:"2miners_url"`
	RedisHost             string   `yaml:"redis_host"`
	RedisPort             int      `yaml:"redis_port"`
	RedisPassword         string   `yaml:"redis_password"`
	RedisShortLifetime    int      `yaml:"redis_short_lifetime"`
	RedisMidLifetime      int      `yaml:"redis_mid_lifetime"`
	RedisLongLifetime     int      `yaml:"redis_long_lifetime"`
	Factor                float64  `yaml:"factor"`
	EthFactor             float64  `yaml:"ether_factor"`
	GazFactor             float64  `yaml:"gaz_factor"`
	MinerListing          string   `yaml:"miner_listing"`
	APIPort               int      `yaml:"api_port"`
	APIFrontPort          int      `yaml:"api_front_port"`
	APILogFile            string   `yaml:"api_log_file"`
	APIUsername           string   `yaml:"api_username"`
	APIPassword           string   `yaml:"api_password"`
	APIAdress             string   `yaml:"api_adress"`
	AdressFilePath        string   `yaml:"adress_file_path"`
	LockPath              string   `yaml:"lock_path"`
	CoinList              []string `yaml:"coin_list"`
	Adresses
}

func LoadYamlConfig(ConfigFilePath string) {
	t := Config{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	log2miner.Error(err)
	err = yaml.Unmarshal(data, &t)
	log2miner.Error(err)
	Cfg = &t
	data, err = ioutil.ReadFile(Cfg.AdressFilePath)
	log2miner.Error(err)
	a := Adresses{}
	err = yaml.Unmarshal(data, &a)
	log2miner.Error(err)
	Cfg.Adress = a.Adress
}
