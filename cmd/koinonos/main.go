package main

import (
	"fmt"
	"github.com/kibis-is/koinonos/internal/constants"
	"github.com/kibis-is/koinonos/internal/routes"
	"github.com/kibis-is/koinonos/internal/types"
	"github.com/kibis-is/koinonos/internal/utilities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"os"
)

var Version string

func main() {
	var err error

	clientSecret := os.Getenv("CLIENT_SECRET")
	debug := utilities.GetEnvWithDefault("DEBUG", "false") == "true"
	level := slog.LevelError

	// if debug is enabled, set the log level to debug
	if debug {
		level = slog.LevelDebug
	}

	slog.SetLogLoggerLevel(level)

	if clientSecret == "" {
		slog.Debug("client secret not set, generating a new one")

		clientSecret, err = utilities.CreateToken(32)

		if err != nil {
			slog.Error(fmt.Sprintf("failed to create client secret: %v", err))

			os.Exit(1)
		}
	}

	config := types.Config{
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

	// display message
	fmt.Println(`
██╗  ██╗ ██████╗ ██╗███╗   ██╗ ██████╗ ███╗   ██╗ ██████╗ ███████╗
██║ ██╔╝██╔═══██╗██║████╗  ██║██╔═══██╗████╗  ██║██╔═══██╗██╔════╝
█████╔╝ ██║   ██║██║██╔██╗ ██║██║   ██║██╔██╗ ██║██║   ██║███████╗
██╔═██╗ ██║   ██║██║██║╚██╗██║██║   ██║██║╚██╗██║██║   ██║╚════██║
██║  ██╗╚██████╔╝██║██║ ╚████║╚██████╔╝██║ ╚████║╚██████╔╝███████║  v1.0.0

An AVM node companion tool that allows Kibisis and other compatible wallets to interact with your AVM node.
https://koinonos.kibis.is
__________________________________________________________________

`)

	// start the server
	err = e.Start(fmt.Sprintf(":%s", utilities.GetEnvWithDefault("PORT", "8080")))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to start server: %v", err))

		os.Exit(1)
	}
}
