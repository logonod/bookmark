package app

import (
	"context"
	"github.com/sirupsen/logrus"

	"e.coding.net/logonod/note-ms-server/cache"
	"e.coding.net/logonod/note-ms-server/db"
	"e.coding.net/logonod/note-ms-server/mq"
)

type App struct {
	Config   *Config
	Database *db.Database
	Cache    *cache.Cache
	Mq       *mq.Mq
}

func (a *App) NewContext() *Context {
	return &Context{
		Logger:   logrus.StandardLogger(),
		Database: a.Database,
		Cache:    a.Cache,
		Mq:       a.Mq,
		Config:   a.Config,
	}
}

func New() (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	dbConfig, err := db.InitConfig()
	if err != nil {
		return nil, err
	}

	cacheConfig, err := cache.InitConfig()
	if err != nil {
		return nil, err
	}

	mqConfig, err := mq.InitConfig()
	if err != nil {
		return nil, err
	}

	app.Database, err = db.New(dbConfig)
	if err != nil {
		return nil, err
	}

	app.Cache, err = cache.New(cacheConfig)
	if err != nil {
		return nil, err
	}

	app.Mq, err = mq.New(mqConfig)
	if err != nil {
		return nil, err
	}

	return app, err
}

func (a *App) Close() error {
	var err error

	err = a.Database.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	err = a.Cache.Close()
	if err != nil {
		return err
	}

	err = a.Mq.Close()
	if err != nil {
		return err
	}

	return nil
}
