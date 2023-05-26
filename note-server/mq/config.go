package mq

import (
  "github.com/spf13/viper"
)

type Config struct {
  Addr string
}

func InitConfig() (*Config, error) {
  config := &Config{
    Addr: viper.GetString("Mq.Addr"),
  }
  if config.Addr == "" {
    config.Addr = "localhost:11300"
  }

  return config, nil
}
