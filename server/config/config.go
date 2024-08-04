package config

import (
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	Port int        `mapstructure:"port"`
	Auth AuthConfig `mapstructure:"auth"`
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
	viper.AddConfigPath("server")
	viper.SetConfigType("toml")
}

func initDefault() {
	viper.SetDefault("auth.enable_tls", false)
	viper.SetDefault("port", 50051)
}
