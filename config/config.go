package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
)

type Config struct {
	App   `mapstructure:"app"`
	HTTP  `mapstructure:"http"`
	Mysql `mapstructure:"mysql"`
	Auth  `mapstructure:"auth"`
}

type App struct {
	Name string `mapstructure:"name"`
}

type HTTP struct {
	Port string `mapstructure:"port"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
}

type Auth struct {
	Secret      string `mapstructure:"secret`
	ExpiryHours int    `mapstructure:"expiryHours"`
}

var configMode = "dev"

var Module = fx.Module("config", fx.Provide(NewConfig))

func NewConfig() (*Config, error) {
	cfg := new(Config)

	viper.SetConfigName(configMode)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error read config file\n: %v", err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("error unmarshal config file\n: %v", err)
	}

	return cfg, nil
}
