package app

import (
	"fmt"
	"net/http"

	log "github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/lincolnzhou/check-in/backend/app/redis"
	"github.com/lincolnzhou/check-in/backend/app/util"
	"github.com/lincolnzhou/check-in/backend/conf"
)

// InitRouter init router
func InitRouter() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		err := redis.Incr("hit:index")
		if err != nil {
			log.Infof("router / redis inrc error: %s", err.Error())
		}

		return c.String(http.StatusOK, "check in system")
	})

	e.GET("/api/check_count", func(c echo.Context) error {
		cacheKey := fmt.Sprintf("check_in:%d", 1)
		count, err := redis.BitCount(cacheKey)
		if err != nil {
			log.Infof("router / redis bitcount error: %s", err.Error())
		}
		return SetJson(c, 0, count, "")
	})

	e.POST("/api/check", func(c echo.Context) error {
		day := util.TimeDayDiff(util.TimeNow(), conf.ConfigData.StartTime)
		if day >= 0 {
			cacheKey := fmt.Sprintf("check_in:%d", 1)
			err := redis.SetBit(cacheKey, day, true)
			if err != nil {
				log.Infof("router / redis setbit error: %s", err.Error())
			}
		}

		return c.String(http.StatusOK, "checked")
	})

	err := e.Start(conf.ConfigData.ApiListen)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func SetJson(c echo.Context, code int, data interface{}, msg string) error {
	jsonData := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	return c.JSON(http.StatusOK, jsonData)
}
