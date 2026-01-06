package parser

import "strings"

func Parse(input string) (string, string) {
	parts := strings.Fields(input)

	if len(parts) == 0 {
		return "", ""
	}

	cmd := parts[0]
	args := strings.Join(parts[1:], " ")
	return cmd, args
}
