package readers

import (
	"os"
	"strings"
)

func ReadFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func ReadLines(filename string) ([]string, error) {
	str, err := ReadFileToString(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(str, "\n"), nil
}
