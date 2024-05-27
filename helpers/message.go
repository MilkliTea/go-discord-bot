package helpers

import "strings"

func ContainsCaptcha(content string) bool {
	return strings.Contains(content, "captcha")
}

func ContainsSpentCaught(content string) bool {
	return strings.Contains(content, "caught")
}

func ContainsInventory(content string) bool {
	return strings.Contains(content, "Inventory")
}
