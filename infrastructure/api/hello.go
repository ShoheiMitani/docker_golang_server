package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func SayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Beautiful World!")
}
