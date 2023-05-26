package db

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	*mongo.Client
}

func New(config *Config) (*Database, error) {
	clientOptions := options.Client().ApplyURI(config.DatabaseURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	// 检测连接是否建立成功
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	return &Database{client}, nil
}
