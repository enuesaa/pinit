package serve

import (
	"github.com/google/uuid"
)

func CreateId() string {
	return uuid.New().String()
}
