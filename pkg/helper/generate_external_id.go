package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateExternalID(input string) string {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return ""
	}

	numStr := parts[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return ""
	}

	nextNum := num + 1
	nextID := fmt.Sprintf("I-%04d", nextNum)
	return nextID
}
