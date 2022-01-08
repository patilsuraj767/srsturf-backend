package handler

import (
	"time"

	"github.com/labstack/echo/v4"
)

func Status(c echo.Context) error {
	return c.JSON(200, echo.Map{"Status": "Application is running", "Server-date": time.Now()})
}
