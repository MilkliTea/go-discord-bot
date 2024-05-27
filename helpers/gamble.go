package helpers

import (
	"strings"
)

func GambleWin(content string) bool {

	return !strings.Contains(content, "and you lost it all... :c")
}
