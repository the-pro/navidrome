package db_storm

import (
	"os"
	"path"
	"sync"

	"github.com/asdine/storm"
	"github.com/cloudsonic/sonic-server/conf"
	"github.com/cloudsonic/sonic-server/log"
)

var (
	_dbInstance *storm.DB
	once        sync.Once
)

func Db() *storm.DB {
	once.Do(func() {
		err := os.MkdirAll(conf.Sonic.DbPath, 0700)
		if err != nil {
			panic(err)
		}
		dbPath := path.Join(conf.Sonic.DbPath, "storm.db")
		instance, err := storm.Open(dbPath)
		log.Debug("Opening Storm DB from: " + dbPath)
		if err != nil {
			panic(err)
		}
		_dbInstance = instance
	})
	return _dbInstance
}