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
	TokenWatcher          []string `yaml:"token_watcher"`
	UrlTokenWatcher       []string `yaml:"url_token_watcher"`

	APIAdress       string   `yaml:"api_adress"`
	APIPort         int      `yaml:"api_port"`
	CoinList        []string `yaml:"coin_list"`
	FluxNodeAPIURL  string   `yaml:"flux_node_api_url"`
	Taddr           string   `yaml:"taddr"`
	Zelid           string   `yaml:"zelid"`
	GrafanaToken    string   `yaml:"grafana_api_token"`
	GrafanaUser     string   `yaml:"grafana_user"`
	GrafanaPassword string   `yaml:"grafana_password"`
	GrafanaUrl      string   `yaml:"grafana_url"`
	StreamRAddr     string   `yaml:"streamr_address"`

	ChiaDBPoolUser string `yaml:"chia_db_pool_user"`
	ChiaDBPoolPass string `yaml:"chia_db_pool_pass"`
	ChiaDBPoolHost string `yaml:"chia_db_pool_host"`
	ChiaDBPoolPort string `yaml:"chia_db_pool_port"`
	ChiaDBPoolName string `yaml:"chia_db_pool_name"`
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
