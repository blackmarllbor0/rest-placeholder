package config

import "github.com/spf13/viper"

type IConfig interface {
	GetServerPort() string
}

type Config struct {
	viper *viper.Viper
}

func NewConfig() *Config {
	cfg := Config{}
	cfg.viper = viper.New()
	cfg.viper.AddConfigPath("config")
	cfg.viper.SetConfigName("config")

	return &cfg
}

func (c Config) ReadConfig() error {
	return c.viper.ReadInConfig()
}

func (c Config) GetServerPort() string {
	return c.viper.GetString("server.port")
}
