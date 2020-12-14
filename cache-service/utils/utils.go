package utils

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func GenerateUUID() string {

	uu1, _ := uuid.NewV4()

	return fmt.Sprintf("%s", uu1)
}
