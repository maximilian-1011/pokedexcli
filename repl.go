package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var result []string
	textStriped := strings.TrimSpace(text)
	if textStriped == "" {
		return []string{}
	}
	textLow := strings.ToLower(textStriped)
	result = strings.Split(textLow, " ")
	return result
}
