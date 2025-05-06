package api

type VersionsNodeResponse struct {
	Connected   bool   `json:"connected"`
	GenesisHash string `json:"genesisHash,omitempty"`
	GenesisID   string `json:"genesisID,omitempty"`
}
