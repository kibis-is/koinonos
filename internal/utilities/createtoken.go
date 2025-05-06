package utilities

import (
	"crypto/rand"
	"encoding/hex"
)

// CreateToken generates a random token of the specified length and returns it as a hexadecimal string.
//
// Parameters:
//   - length: The length of the token.
//
// Returns:
//   - A random token of the specified length.
func CreateToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
