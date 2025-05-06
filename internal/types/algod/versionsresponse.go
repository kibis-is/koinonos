package algod

type Build struct {
	Branch      string `json:"branch"`
	BuildNumber int64  `json:"build_number"`
	Channel     string `json:"channel"`
	CommitHash  string `json:"commit_hash"`
	Major       int64  `json:"major"`
	Minor       int64  `json:"minor"`
}

type VersionsResponse struct {
	Build       Build    `json:"build"`
	GenesisHash string   `json:"genesis_hash_b64"`
	GenesisID   string   `json:"genesis_id"`
	Versions    []string `json:"versions"`
}
