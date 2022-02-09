package redis

import (
	"2miner-monitoring/config"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var (
	Clis = make(map[int]*redis.Client, 2)
)

func InitRedis() {

	connect(0)
	go func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				if err := Clis[0].Ping(); err.Err() != nil {
					fmt.Println(err.Err())
					connect(0)
				}
			}
		}
	}()
}
func connect(DbNum int) {
	fmt.Println(config.Cfg.RedisHost, config.Cfg.RedisPort)
	Clis[DbNum] = redis.NewClient(&redis.Options{
		Addr:        config.Cfg.RedisHost + ":" + strconv.Itoa(config.Cfg.RedisPort),
		Password:    config.Cfg.RedisPassword, // no password set
		DB:          DbNum,
		IdleTimeout: 10 * time.Second,
		// use default DB
	})

}

func WriteToRedis(DbNum int, key, value, lifetime string) {
	if lifetime == "long" {
		Clis[DbNum].Set(key, value, time.Duration(config.Cfg.RedisLongLifetime)*time.Second).Err() //TODO: Ajouter du logging et gestion d'erreur

	} else if lifetime == "mid" {
		Clis[DbNum].Set(key, value, time.Duration(config.Cfg.RedisMidLifetime)*time.Second).Err() //TODO: Ajouter du logging et gestion d'erreur
	} else {
		Clis[DbNum].Set(key, value, time.Duration(config.Cfg.RedisShortLifetime)*time.Second).Err() //TODO: Ajouter du logging et gestion d'erreur
	}
}

func GetFromToRedis(DbNum int, key string) string {
	val, err := Clis[DbNum].Get(key).Result()
	if err == redis.Nil || err != nil {
		return ""
	}
	return val
}
