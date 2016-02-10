package controller

import (
	"github.com/garyburd/redigo/redis"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

var (
	//Render :render data var
	Render *render.Render
	//S :mgo session
	S            *mgo.Session
	databaseName = "oj"
	//RedisPool :Redis connection pool
	RedisPool *redis.Pool
)

// Init :init controller methods
func init() {
	Render = render.New(render.Options{
		Directory: "templates",
		Layout:    "layout",
	})
	RedisPool = newRedisPool()
}
