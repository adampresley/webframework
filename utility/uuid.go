package utility

import (
	"strings"

	"github.com/nu7hatch/gouuid"
)

/*
GenerateUUID creates a UUID
*/
func GenerateUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

/*
GenerateUUIDWithoutDashes creates a UUID without dashes
*/
func GenerateUUIDWithoutDashes() string {
	uuid, _ := uuid.NewV4()
	stringVersion := strings.Replace(uuid.String(), "-", "", -1)
	return stringVersion
}
