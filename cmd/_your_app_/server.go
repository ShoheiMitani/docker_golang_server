package server

import (
	"docker_golang_server/infrastructure/api"

	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()
	e.GET("/", api.SayHello)
	e.Logger.Fatal(e.Start(":9090"))
}
