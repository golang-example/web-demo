package redis

import (
	. "web-demo/log"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redisPool *redis.Pool
)

const (
	maxIdle     = 500
	idleTimeout = 0 * time.Second
	wait        = true
)

//Init
//eg: RedisInit("127.0.0.1:6379", 0, "pwd", 8)
func RedisInit(server string, db int, password string, maxConn int) {
	maxActive := maxConn

	//Make a pool object
	redisPool = &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: maxIdle,

		// Maximum number of connections allocated by the pool at a given time.
		// When zero, there is no limit on the number of connections in the pool.
		MaxActive: maxActive,

		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: idleTimeout,

		// If Wait is true and the pool is at the MaxActive limit, then Get() waits
		// for a connection to be returned to the pool before returning.
		Wait: wait,

		// Dial is an application supplied function for creating and configuring a
		// connection.
		//
		// The connection returned from Dial must not be in a special state
		// (subscribed to pubsub channel, transaction started, ...).
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		// TestOnBorrow is an optional application supplied function for checking
		// the health of an idle connection before the connection is used again by
		// the application. Argument t is the time that the connection was returned
		// to the pool. If the function returns an error, then the connection is
		// closed.
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	Log.Info("redis[%v %v] init ok", server, db)
}

// ----
// KeyValue
// ----
func kSet(key, value string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("SET", key, value)
	return redis.String(r, err)
}

func kGet(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("GET", key)
	return redis.String(r, err)
}

func kMGet(keys []interface{}) ([]string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("MGET", keys...)
	return redis.Strings(r, err)
}

func kDel(key string) (uint64, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("DEL", key)
	return redis.Uint64(r, err)
}

func expire(key string, seconds int)(int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("EXPIRE", key, seconds);
	return redis.Int(r, err)
}

func kSetex(key string, seconds int, value string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("SETEX", key, seconds, value)
	return redis.String(r, err)
}

func kSetnx(key string, value string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("SETNX", key, value)
	return redis.Int(r, err)
}

// ----
//Hash
// ----
func hSet(key, field, data string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("HSET", key, field, data)
	return redis.Int(r, err)
}

func hMset(key string, keyfields interface{}) (string, error) {
	args := redis.Args{}.Add(key).AddFlat(keyfields)
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("HMSET", args...)
	return redis.String(r, err)
}

func hDel(key string, fields interface{}) (int, error) {
	args := redis.Args{}.Add(key).AddFlat(fields)
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("HDEL", args...)
	return redis.Int(r, err)
}

func hGet(key, field string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("HGET", key, field)
	return redis.String(r, err)
}

func hMget(key string, fields []string) ([]string, error) {
	args := redis.Args{}.Add(key).AddFlat(fields)
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("HMGET", args...)
	return redis.Strings(r, err)
}

func hLen(key string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("HLEN", key)
	return redis.Int(r, err)
}

func hGetAll(key string) (map[string]string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("HGETALL", key)
	return redis.StringMap(r, err)
}

// ----
// Set
// ----
func sAdd(key, data string) (uint64, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("SADD", key, data)
	return redis.Uint64(r, err)
}

func sCard(key string) (uint64, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("SCARD", key)
	return redis.Uint64(r, err)
}

//exist key
func isExistKey(key string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	num, err := conn.Do("EXISTS", key)
	return redis.Int(num, err)
}

//list
func lPushOneStrVal(key, value string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	num, err := conn.Do("LPUSH", key, value)
	return redis.Int(num, err)
}

func lPUshList(key string, value []string) (int, error) {
	args := redis.Args{}.Add(key).AddFlat(value)
	conn := redisPool.Get()
	defer conn.Close()
	num, err := conn.Do("LPUSH", args...)
	return redis.Int(num, err)
}

func rPUshList(key string, value []string) (int, error) {
	args := redis.Args{}.Add(key).AddFlat(value)
	conn := redisPool.Get()
	defer conn.Close()
	num, err := conn.Do("RPUSH", args...)
	return redis.Int(num, err)
}

func rPushOneStrVal(key, value interface{}) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	num, err := conn.Do("RPUSH", key, value)
	return redis.Int(num, err)
}

func lPop(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	value, err := conn.Do("LPOP", key)
	return redis.String(value, err)
}

func rPop(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	value, err := conn.Do("RPOP", key)
	return redis.String(value, err)
}

func lLen(key string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	len, err := conn.Do("LLEN", key)
	return redis.Int(len, err)
}

func lRange(key string, start, end int) ([]string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("LRANGE", key, start, end)
	return redis.Strings(r, err)
}

func lRem(key, value string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("LREM", key, 0, value)
	return redis.Int(r, err)
}

func lIndex(key string, index int) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	v, err := conn.Do("LINDEX", key, index)
	return redis.String(v, err)
}

func lTrim(key string, start, end int)(interface{}, error){
	conn := redisPool.Get()
	defer conn.Close()
	v, err := conn.Do("LTRIM", key, start, end)
	return redis.Values(v, err)
}

//key ttl
func kTtl(key string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	v, err := conn.Do("TTL", key)
	return redis.Int(v, err)
}

//Zset
func zAdd(key, data string, score int) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("ZADD", key, score, data)
	return redis.Int(r, err)
}

func zCard(key string) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("ZCARD", key)
	return redis.Int(r, err)
}

func zRemRangeByScore(key string, score1, score2 int) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	r, err := conn.Do("ZREMRANGEBYSCORE", key, score1, score2)
	return redis.Int(r, err)
}
