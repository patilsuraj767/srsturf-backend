package turf

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patilsuraj767/turf/turf/config"
	"github.com/patilsuraj767/turf/turf/db"
)

func StartServer() {
	config.InitConfig()
	db.InitDatabase()
	e := echo.New()
	e.Debug = true
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	SetupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
