package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func SetJson(c echo.Context, code int, data interface{}, msg string) error {
	jsonData := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	return c.JSON(http.StatusOK, jsonData)
}
