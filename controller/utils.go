package controller

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
)

func getMongoS() *mgo.Session {
	if S == nil {
		var err error
		S, err = mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
	}
	return S.Clone()
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}
