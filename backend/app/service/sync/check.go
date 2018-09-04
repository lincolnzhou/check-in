package sync

import "github.com/lincolnzhou/check-in/backend/conf"

type Checker struct {
	ring *Ring
	conf conf.Config
}

type Check struct {
	ID       int32
	DateTime int64
}

func NewSyncChecker(conf *conf.Config) (c *Checker) {
	c = &Checker{}
	return
}
