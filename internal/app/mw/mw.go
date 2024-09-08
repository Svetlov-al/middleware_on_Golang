package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"strings"
)

const roleAdmin = "admin"

func RoleCheckMW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		val := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(val, roleAdmin) {
			log.Info().Msg("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
