package main

import (
	"fmt"
	"github.com/kibis-is/koinonos/internal/constants"
	"github.com/kibis-is/koinonos/internal/routes"
	commontypes "github.com/kibis-is/koinonos/internal/types/common"
	cryptographyutilities "github.com/kibis-is/koinonos/internal/utilities/cryptography"
	osutilities "github.com/kibis-is/koinonos/internal/utilities/os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"os"
)

var Version string

func main() {
	var err error

	clientSecret := os.Getenv("CLIENT_SECRET")
	debug := osutilities.GetEnvWithDefault("DEBUG", "false") == "true"
	level := slog.LevelError

	// if debug is enabled, set the log level to debug
	if debug {
		level = slog.LevelDebug
	}

	slog.SetLogLoggerLevel(level)

	if clientSecret == "" {
		slog.Debug("client secret not set, generating a new one")

		clientSecret, err = cryptographyutilities.CreateToken(32)

		if err != nil {
			slog.Error(fmt.Sprintf("failed to create client secret: %v", err))

			os.Exit(1)
		}
	}

	config := commontypes.Config{
		ClientSecret: clientSecret,
		Debug:        debug,
		Version:      Version,
	}

	e := echo.New()

	e.HideBanner = true

	// middlewares
	e.Use(middleware.Logger())

	// routes
	// /versions
	e.GET(constants.VersionsPath, routes.NewGetVersionsRoute(config))

	// display a welcome message
	fmt.Println(constants.WelcomeMessage)

	// start the server
	err = e.Start(fmt.Sprintf(":%s", osutilities.GetEnvWithDefault("PORT", "3000")))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to start server: %v", err))

		os.Exit(1)
	}
}
