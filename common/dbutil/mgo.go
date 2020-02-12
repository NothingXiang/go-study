/*
 *数据库连接相关工具
 */
package dbutil

import (
	"log"
	"sync"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	once    sync.Once
)

func NewSession() *mgo.Session {
	// make sure just connect a session once
	once.Do(func() {
		mgoUrl := viper.GetString("mongo.url")

		var err error
		if session, err = mgo.Dial(mgoUrl); err != nil {
			log.Panicf("[NewSession] dial mongo failed:%v", err)
		}

	})
	return session.Clone()
}

func NewMgoDB() *mgo.Database {
	return NewSession().DB(viper.GetString("mongo.dbname"))
}
