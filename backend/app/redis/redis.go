package redis

import (
	"github.com/lincolnzhou/check-in/backend/conf"
)

// Set redis set
func Set(key string, value string) error {
	conn := conf.ConfigData.Redis.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", key, value)
	if err != nil {
		return err
	}

	return nil
}

// Incr redis increment
func Incr(key string) error {
	conn := conf.ConfigData.Redis.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("incr", key)
	if err != nil {
		return err
	}

	return nil
}
