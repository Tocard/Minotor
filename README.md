# Minotor
This go application is an upgrade of the 2miner-monitoring.
His goal is to be able to crawl more data than the original 2miner-monitoring.

If you want to monitore only yourself, you should considerate 2miner-monitoring.

This application harvest data from 2miners API (https://apidoc.2miners.com/#/)

It send the data to an elasticsearch and on my case grafana is used to retrieves visualisation

## Important Url

API         => https://load01.ether-source.fr:8999

Grafana     => https://load01.ether-source.fr

## How it works

### The Backend

Gin is used to create a http server with many handler
Gocron is used to proc all call at a given time to harvest data  

### The FrontEnd

There is a grafana to display everything. No registration is needed for viewer only.
If you wanna do your own dashboard/alert, you need to register.

### Service connected

- 2miners
- Hiveos 
- Etherscan
- Hashrate.no
- Coingecko

#### 2miners

To un/subscribe to the monitoring, you have only to do an api call
````go
	server.GET("/subscribe/:wallet", handlers.SuscribeWallet)
	server.GET("/unsubscribe/:wallet", handlers.UnSuscribeWallet)
````

After that, the server will collectd data from the given wallet


#### HiveOS

To have the Hiveos, the client have to allow **minotor** permission to the user **minotor_LFDM**

#### The Others

the others services are called on demand or periodically to populate the data

## How to run it

### Config file

You have ton configure the yaml files(adresses, cards and config). All field is currently mandatory.


### Install requirement

GO >= 1.17, Elasticsearch (cluster is nice to have), grafana, redis.
Haproxy is a must


### Run the binary

The binary is the brain of the process, everything is handle by him. Technically you don't need ES + REDIS + GRAFANA, that's just a must and tooling.

### TroobleShooting
