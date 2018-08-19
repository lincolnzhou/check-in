package app

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lincolnzhou/check-in/backend/app/redis"
	"github.com/lincolnzhou/check-in/backend/conf"
)

func InitRouter() {
	// 初始化HTTP服务
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		redis.Incr("hit:index")

		return c.String(http.StatusOK, "check in system")
	})

	e.Start(conf.ConfigData.ApiListen)
}
