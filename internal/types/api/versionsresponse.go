package api

type VersionsResponse struct {
	App  VersionsAppResponse  `json:"app"`
	Node VersionsNodeResponse `json:"node"`
}
