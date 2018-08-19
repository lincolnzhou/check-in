package app

import (
	"net/http"

	log "github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/lincolnzhou/check-in/backend/app/redis"
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

	e.Start(conf.ConfigData.ApiListen)
}
