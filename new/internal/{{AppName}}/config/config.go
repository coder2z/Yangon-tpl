package config

import (
	"fmt"
	"github.com/spf13/viper"
	"{{ProjectName}}/internal/{{AppName}}/config/server"
	"{{ProjectName}}/pkg/client/database"
	"{{ProjectName}}/pkg/log"
)

const (
	// DefaultConfigurationName is the default name of configuration
	defaultConfigurationName = "config"

	// DefaultConfigurationPath the default location of the configuration file
	defaultConfigurationPath = "./config"
)

type Cfg struct {
	Server *server.Options
	Mysql  *database.Options
	Log    *log.Options
}

func cfg() *Cfg {
	return &Cfg{
		Server: server.NewServerOptions(),
		Mysql:  database.NewDatabaseOptions(),
		Log:    log.NewLogOptions(),
	}
}

func TryLoadFromDisk() (*Cfg, error) {
	viper.SetConfigName(defaultConfigurationName)
	viper.AddConfigPath(defaultConfigurationPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		} else {
			return nil, fmt.Errorf("error parsing configuration file %s", err)
		}
	}

	conf := cfg()

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
