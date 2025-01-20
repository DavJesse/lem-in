package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ValidContent(filename string) ([]string, error) {
	fileContent, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	defer fileContent.Close()

	validContent := []string{}
	scanner := bufio.NewScanner(fileContent)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		validContent = append(validContent, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return validContent, nil
}

func CheckArgs() {
	if len(os.Args) != 2 {
		log.Fatal("ERROR: invalid data format, please provide a file name")
	}
}
