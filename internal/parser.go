package internal

import (
	"fmt"
	"strings"
)

func LoadFromString(data string) (MapOfMaps, error) {
	lines := strings.Split(data, "\n")
	cleanLines := make([]string, 0)
	parsedData := make(MapOfMaps)
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		if !strings.HasPrefix(line, "[") {
			line = strings.ReplaceAll(line, " ", "")
		}

		cleanLines = append(cleanLines, line)

	}

	for _, line := range cleanLines {
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.Trim(line, "[]")

			parsedData[currentSection] = make(map[string]string)

		} else if currentSection != "" {
			trimmedLine := strings.Split(line, "=")

			if len(trimmedLine) != 2 || trimmedLine[0] == "" || trimmedLine[1] == "" {
				return nil, fmt.Errorf("invalid line format")
			}

			key := trimmedLine[0]
			value := trimmedLine[1]

			parsedData[currentSection][key] = value
		} else {
			return nil, fmt.Errorf("key-value pair found outside of a section")
		}

	}

	for _, kv := range parsedData {
		fmt.Println(kv)
	}

	return parsedData, nil
}
