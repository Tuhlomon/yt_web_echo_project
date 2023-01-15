package mw

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if strings.EqualFold(ctx.Request().Header.Get("User-Role"), roleAdmin) {
			log.Println("red button user detected")
		}
		return next(ctx)
	}
}
