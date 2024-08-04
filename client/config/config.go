package config

import (
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	ControllerAddr string     `mapstructure:"controller_addr"`
	Auth           AuthConfig `mapstructure:"auth"`
}

type AuthConfig struct {
	EnableTLS  bool   `mapstructure:"enable_tls"`
	PublicKey  string `mapstructure:"public_key"`
	PrivateKey string `mapstructure:"private_key"`
	CACert     string `mapstructure:"ca_cert"`
}

var (
	config   GlobalConfig
	isLoaded bool
)

func Config() *GlobalConfig {
	if !isLoaded {
		initViper()
		load()
	}
	return &config
}

func load() {
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: " + err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}

	isLoaded = true
}

func initViper() {
	initDefault()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("client")
	viper.SetConfigType("toml")
}

func initDefault() {
	viper.SetDefault("auth.enable_tls", false)
	viper.SetDefault("connection.controller_addr", "localhost:50051")
}
