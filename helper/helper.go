package helper

import (
	"os"
	"strings"
)

func GetInputLines(path string) []string {
	inputBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	return lines
}
