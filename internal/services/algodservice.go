package services

import (
	"fmt"
	"github.com/kibis-is/koinonos/internal/constants"
	algodtypes "github.com/kibis-is/koinonos/internal/types/algod"
	apiutilities "github.com/kibis-is/koinonos/internal/utilities/api"
	"log/slog"
	"net/http"
)

type AlgodService struct {
}

func NewAlgodService() *AlgodService {
	return &AlgodService{}
}

/*
 * public functions
 */

func (s *AlgodService) GetVersions() (*algodtypes.VersionsResponse, error) {
	var result *algodtypes.VersionsResponse

	response, err := http.Get(fmt.Sprintf("%s/versions", constants.AlgodURL))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get algod versions: %v", err))

		return nil, err
	}

	defer response.Body.Close()

	if err = apiutilities.ParseResponseBody(response.Body, &result); err != nil {
		slog.Error(fmt.Sprintf("failed to parse /versions response: %v", err))

		return nil, err
	}

	return result, nil
}
