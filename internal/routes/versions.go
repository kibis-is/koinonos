package routes

import (
	"github.com/kibis-is/koinonos/internal/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewGetVersionsRoute(config types.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, types.VersionsResponse{
			Version: config.Version,
		})
	}
}
