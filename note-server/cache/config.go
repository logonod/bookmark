package cache

import (
  "github.com/spf13/viper"
)

type Config struct {
  Addr     string
  Password string
  DB       int
}

func InitConfig() (*Config, error) {
  config := &Config{
    Addr:     viper.GetString("Cache.Addr"),
    Password: viper.GetString("Cache.Password"),
    DB:       viper.GetInt("Cache.DB"),
  }
  if config.Addr == "" {
    config.Addr = "localhost:6379"
  }

  return config, nil
}
