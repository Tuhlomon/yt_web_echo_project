package app

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"github.com/Tuhlomon/yt_web_echo_project/internal/app/endpoint"
	"github.com/Tuhlomon/yt_web_echo_project/internal/app/mw"
	"github.com/Tuhlomon/yt_web_echo_project/internal/app/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.GET("/status", a.e.Status, mw.RoleCheck)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running...")

	err := a.echo.Start(":8080")

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
