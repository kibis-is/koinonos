package routes

import (
	"github.com/kibis-is/koinonos/internal/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewGetVersionsRoute(version string) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, types.VersionsResponse{
			Version: version,
		})
	}
}
