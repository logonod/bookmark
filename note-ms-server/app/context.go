package app

import (
	"github.com/sirupsen/logrus"

	"e.coding.net/logonod/note-ms-server/cache"
	"e.coding.net/logonod/note-ms-server/db"
	"e.coding.net/logonod/note-ms-server/model"
	"e.coding.net/logonod/note-ms-server/mq"
)

type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	Database      *db.Database
	Cache         *cache.Cache
	Mq            *mq.Mq
	User          *model.User
	Config        *Config
}

func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}

func (ctx *Context) WithUser(user *model.User) *Context {
	ret := *ctx
	ret.User = user
	return &ret
}
