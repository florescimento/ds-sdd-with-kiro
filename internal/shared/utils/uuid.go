package utils

import "github.com/google/uuid"

// GenerateUUID generates a new UUIDv4
func GenerateUUID() string {
	return uuid.New().String()
}

// ValidateUUID checks if a string is a valid UUID
func ValidateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
