package redis

import (
	redigo "github.com/gomodule/redigo/redis"
	"github.com/lincolnzhou/check-in/backend/conf"
)

// Set redis set
func Set(key string, value string) error {
	conn := conf.ConfigData.Redis.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// Incr redis increment
func Incr(key string) error {
	conn := conf.ConfigData.Redis.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("INCR", key)
	if err != nil {
		return err
	}

	return nil
}

// SetBit redis setbit
func SetBit(key string, offset int, value bool) error {
	conn := conf.ConfigData.Redis.Pool.Get()
	defer conn.Close()

	v := 0
	if value {
		v = 1
	}

	_, err := conn.Do("SETBIT", key, offset, v)
	if err != nil {
		return err
	}

	return nil
}

// BitCount redis bitcount
func BitCount(key string) (int, error) {
	conn := conf.ConfigData.Redis.Pool.Get()
	defer conn.Close()

	count, err := redigo.Int(conn.Do("BITCOUNT", key))
	if err != nil {
		return 0, err
	}

	return count, nil
}
