package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	generatedUUID := uuid.New()
	return generatedUUID.String()
}

func ParseUUIDFromString(s string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(s)
	return parsed, err
}
