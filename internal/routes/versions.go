package routes

import (
	"github.com/kibis-is/koinonos/internal/services"
	apitypes "github.com/kibis-is/koinonos/internal/types/api"
	commontypes "github.com/kibis-is/koinonos/internal/types/common"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewGetVersionsRoute(config commontypes.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		algodService := services.NewAlgodService()
		versions, _ := algodService.GetVersions()
		node := apitypes.VersionsNodeResponse{
			Connected: false,
		}

		if versions != nil {
			node.Connected = true
			node.GenesisHash = versions.GenesisHash
			node.GenesisID = versions.GenesisID
		}

		return c.JSON(http.StatusOK, apitypes.VersionsResponse{
			App: apitypes.VersionsAppResponse{
				Version: config.Version,
			},
			Node: node,
		})
	}
}
