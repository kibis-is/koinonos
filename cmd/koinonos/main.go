package main

import (
	"fmt"
	"github.com/kibis-is/koinonos/internal/constants"
	"github.com/kibis-is/koinonos/internal/routes"
	"github.com/kibis-is/koinonos/internal/utilities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"os"
)

var Version string

func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())

	// routes
	// /versions
	e.GET(constants.VersionsPath, routes.NewGetVersionsRoute(Version))

	// start the server
	err := e.Start(fmt.Sprintf(":%s", utilities.GetEnvWithDefault("PORT", "8080")))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to start server: %v", err))

		os.Exit(1)
	}
}
