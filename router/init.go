package router

import (
	"database/sql"

	"github.com/gomodule/redigo/redis"
)

var (
	dbPool    *sql.DB
	CachePool *redis.Pool
)

func Init(db *sql.DB, cache *redis.Pool) {
	dbPool = db
	CachePool = cache
}
