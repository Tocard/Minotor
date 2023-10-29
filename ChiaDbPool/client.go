package ChiaDbPool

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"minotor/config"
)

var (
	Client *gorm.DB
)

func ConnectToDB() error {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", config.Cfg.ChiaDBPoolUser, config.Cfg.ChiaDBPoolPass,
		config.Cfg.ChiaDBPoolName, config.Cfg.ChiaDBPoolHost, config.Cfg.ChiaDBPoolPort)
	var err error
	Client, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
