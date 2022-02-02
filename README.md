# 2miner-monitoring-go
This go application is an upgrade of the 2miner-monitoring.
His goal is to be able to crawl more data than the original 2miner-monitoring.

If you want to monitore only yourslef, you should considerate 2miner-monitoring.

This application harvest data from 2miners API (https://apidoc.2miners.com/#/)

It send the data to an elasticsearch and on my case grafana is used to retrieves visualisation

## How to run it

### Config file

You have ton configure the config.yaml file (default is config/config.yaml). All field is currently mandatory.

```yaml
---
elasticsearch_user: "es_user"
elasticsearch_password: "es_password"
elasticsearch_hosts:  ["host1:port", "host2:port"...]
api_token_etherscan: "etherscan_api_token "
log_level: INFO # CRITICAL => ERROR => WARNING => INFO => DEBUG
ca_path: "ssl/chain.pem"
2miners_url: "https://eth.2miners.com/api"

miner_listing: ADDR #ALL is used to get all miners from 2miners, ADDR is used for adress variable
adress:
  - "0x........................................"
  - "0x........................................"
  - "0x........................................"
  - "0x........................................"

redis_host: 127.0.0.1
redis_port: 6379
redis_password: ""
redis_lifetime: 10

factor: 0.000001
ether_factor: 0.000000000000000001
gas_factor: 0.000000001

```

### Install requirement


### Run the binary


### TroobleShooting
