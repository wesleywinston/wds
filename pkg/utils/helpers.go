package utils

import (
	"fmt"
)

// FormatMessage takes a name and returns a formatted greeting string.
// Note: Function must be capitalized (exported) to be used outside the package.
func FormatMessage(fullName []string) string {
	return fmt.Sprintf("Server up and running! Greeting from utils, %s %s.", fullName[0], fullName[1])
}
