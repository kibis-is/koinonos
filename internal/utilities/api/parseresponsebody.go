package api

import (
	"encoding/json"
	"io"
)

// ParseResponseBody Attempts to parse a JSON HTTP response to the provided interface.
//
// Parameters:
//   - responseBody: The response body.
//   - output: The interface to output to.
//
// Returns:
//   - An error if something went wrong.
func ParseResponseBody(responseBody io.Reader, output any) error {
	body, err := io.ReadAll(responseBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &output)
}
