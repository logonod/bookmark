package api

import "github.com/spf13/viper"

type Config struct {
	// The port to bind the web application server to
	Port int

	// The number of proxies positioned in front of the API. This is used to interpret
	// X-Forwarded-For headers.
	ProxyCount int

	SessionTTL int
}

func InitConfig() (*Config, error) {
	config := &Config{
		Port:       viper.GetInt("Port"),
		ProxyCount: viper.GetInt("ProxyCount"),
		SessionTTL: viper.GetInt("SessionTTL"),
	}
	if config.Port == 0 {
		config.Port = 8080
	}
	if config.SessionTTL == 0 {
		config.SessionTTL = 2592000
	}

	return config, nil
}
