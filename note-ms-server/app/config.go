package app

import (
  "fmt"

  "github.com/spf13/viper"
)

type Config struct {
  // A secret string used for session cookies, passwords, etc.
  SecretKey []byte
  // 标题保存最大长度
  NormSubStringTitleLength int
  // 描述保存最大长度
  NormSubStringDescriptionLength int
  // 全文保存最大长度
  NormSubFullTextLength int
  // 域名最大长度
  NormSubStringDomainLength int
}

func InitConfig() (*Config, error) {
  config := &Config{
    SecretKey:                      []byte(viper.GetString("SecretKey")),
    NormSubStringTitleLength:       viper.GetInt("NormSubStringTitleLength"),
    NormSubStringDescriptionLength: viper.GetInt("NormSubStringDescriptionLength"),
    NormSubFullTextLength:          viper.GetInt("NormSubFullTextLength"),
    NormSubStringDomainLength:      viper.GetInt("NormSubStringDomainLength"),
  }
  if len(config.SecretKey) == 0 {
    return nil, fmt.Errorf("SecretKey must be set")
  }

  return config, nil
}
