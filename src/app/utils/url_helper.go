package utils

import (
	"fmt"
	"strings"
)

func GetScope(version string) string {
	scope := ""
	if strings.TrimSpace(version) != "" {
		scope = fmt.Sprintf("version=%s", version)
	}

	return scope
}
