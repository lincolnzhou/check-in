package app

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/lincolnzhou/check-in/backend/app/controller"
	"github.com/lincolnzhou/check-in/backend/conf"
)

// InitRouter init router
func InitRouter() {
	e := echo.New()
	e.Static("/", "static")
	e.GET("/", controller.Index)
	e.GET("/api/hit_count", controller.HitCount)
	e.GET("/api/check", controller.CheckList)
	e.GET("/api/check_count", controller.CheckCount)
	e.POST("/api/check", controller.CheckPost)

	err := e.Start(conf.ConfigData.ApiListen)
	if err != nil {
		fmt.Println(err.Error())
	}
}
