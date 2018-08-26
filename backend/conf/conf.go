package conf

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
)

type Config struct {
	Version   string
	ApiListen string
	Redis     *Redis
	StartTime string
}

type Redis struct {
	Host     string
	Password string
	Db       int

	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Wait        bool

	ReadTimeout    int
	WriteTimeout   int
	ConnectTimeout int

	Pool *redis.Pool
}

var (
	ConfigData *Config
)

// NewConfig new config
func NewConfig(conf string) (c *Config, err error) {
	var (
		file *os.File
		blob []byte
	)

	c = new(Config)
	if file, err = os.Open(conf); err != nil {
		return
	}
	if blob, err = ioutil.ReadAll(file); err != nil {
		return
	}

	err = toml.Unmarshal(blob, c)
	return
}

func NewPool(cnf *Redis) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     cnf.MaxIdle,
		IdleTimeout: time.Duration(cnf.IdleTimeout) * time.Second,
		MaxActive:   cnf.MaxActive,
		Wait:        cnf.Wait,
		Dial: func() (redis.Conn, error) {
			var opts = []redis.DialOption{
				redis.DialDatabase(cnf.Db),
				redis.DialReadTimeout(time.Duration(cnf.ReadTimeout) * time.Second),
				redis.DialWriteTimeout(time.Duration(cnf.WriteTimeout) * time.Second),
				redis.DialConnectTimeout(time.Duration(cnf.ConnectTimeout) * time.Second),
			}

			if cnf.Password != "" {
				opts = append(opts, redis.DialPassword(cnf.Password))
			}

			c, err := redis.Dial("tcp", cnf.Host, opts...)
			if err != nil {
				return nil, err
			}

			if cnf.Db != 0 {
				_, err = c.Do("SELECT", cnf.Db)
				if err != nil {
					return nil, err
				}
			}

			return c, err
		},
		// PINGs connections that have been idle more than 10 seconds
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Duration(10*time.Second) {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitConfig(conf string) (err error) {
	ConfigData, err = NewConfig(conf)
	if err != nil {
		return err
	}

	if os.Getenv("REDIS_HOST") != "" {
		ConfigData.Redis.Host = os.Getenv("REDIS_HOST")
	}

	ConfigData.Redis.Pool = NewPool(ConfigData.Redis)
	return err
}
