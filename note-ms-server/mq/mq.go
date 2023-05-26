package mq

import (
  beanstalk "github.com/beanstalkd/go-beanstalk"
  "github.com/pkg/errors"
)

type Mq struct {
  *beanstalk.Conn
}

func New(config *Config) (*Mq, error) {
  mq, err := beanstalk.Dial("tcp", config.Addr)
  if err != nil {
    return nil, errors.Wrap(err, "unable to connect to beanstalk")
  }

  return &Mq{mq}, nil
}
