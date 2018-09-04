package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/lincolnzhou/check-in/backend/app/redis"
	"github.com/lincolnzhou/check-in/backend/app/util"
	"github.com/lincolnzhou/check-in/backend/conf"
)

func Index(c echo.Context) error {
	err := redis.Incr("hit:index")
	if err != nil {
		log.Infof("router / redis inrc error: %s", err.Error())
	}

	return c.File("static/index.html")
}

func CheckList(c echo.Context) error {
	day := util.TimeDayDiff(util.TimeNow(), conf.ConfigData.StartTime)
	if day >= 0 {
		cacheKey := fmt.Sprintf("check_in:%d", 1)
		bytes, err := redis.GetBits(cacheKey)
		if err != nil {
			log.Infof("router / redis get error: %s", err.Error())
		}

		startTime, _ := time.Parse("2006-01-02 15:04:05", conf.ConfigData.StartTime)
		ret := map[int64]int{}
		for k, v := range bytes {
			str := fmt.Sprintf("%08b", int(v))
			for m, n := range strings.Split(str, "") {
				if n == "1" {
					curTime := startTime.Unix() + int64((k*8+m-1)*24*3600)
					ret[curTime] = 1
				}
			}
		}

		return SetJson(c, 0, ret, "")
	}

	return SetJson(c, 0, nil, "")
}

func HitCount(c echo.Context) error {
	count, err := redis.Get("hit:index")
	if err != nil {
		log.Infof("router / redis get error: %s", err.Error())
	}

	return SetJson(c, 0, count, "")
}

func CheckCount(c echo.Context) error {
	cacheKey := fmt.Sprintf("check_in:%d", 1)
	count, err := redis.BitCount(cacheKey)
	if err != nil {
		log.Infof("router / redis bitcount error: %s", err.Error())
	}
	return SetJson(c, 0, count, "")
}

func CheckPost(c echo.Context) error {
	day := util.TimeDayDiff(util.TimeNow(), conf.ConfigData.StartTime)
	if day >= 0 {
		cacheKey := fmt.Sprintf("check_in:%d", 1)
		err := redis.SetBit(cacheKey, day, true)
		if err != nil {
			log.Infof("router / redis setbit error: %s", err.Error())
		}
	}

	return SetJson(c, 0, "", "")
}
