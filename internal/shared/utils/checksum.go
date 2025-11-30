package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// CalculateSHA256 calculates the SHA256 checksum of data
func CalculateSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// CalculateSHA256FromReader calculates SHA256 from an io.Reader
func CalculateSHA256FromReader(r io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
