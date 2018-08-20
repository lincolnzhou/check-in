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

	e.POST("/api/check", func(c echo.Context) error {
		day := util.TimeDayDiff(util.TimeNow(), conf.ConfigData.StartTime)
		if day >= 0 {
			cache_key := fmt.Sprintf("check_in:%d", 1)
			err := redis.SetBit(cache_key, day, true)
			if err != nil {
				log.Infof("router / redis setbit error: %s", err.Error())
			}
		}

		return c.String(http.StatusOK, "checked")
	})

	e.Start(conf.ConfigData.ApiListen)
}
