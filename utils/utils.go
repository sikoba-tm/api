package utils

import (
	"github.com/google/uuid"
)

func ParseUUIDFromString(s string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(s)
	return parsed, err
}
