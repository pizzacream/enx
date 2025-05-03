package enx

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const EnvVarPattern = `^\s*([\w.-]+)\s*=\s*(.*)?\s*$`

func Load(filePaths ...string) error {
	filePath := ".env"
	if len(filePaths) > 0 {
		filePath = filePaths[0]
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("dotenv: unable to open file: %w", err)
	}
	defer file.Close()

	envVars, err := parseEnvFile(file)
	if err != nil {
		return fmt.Errorf("dotenv: parse error: %w", err)
	}

	for key, value := range envVars {
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}

	return nil
}

// parseEnvFile reads a file line-by-line and extracts valid environment variables.
func parseEnvFile(file *os.File) (map[string]string, error) {
	scanner := bufio.NewScanner(file)
	pattern, err := regexp.Compile(EnvVarPattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	envVars := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		key, value, ok := parseLine(line, pattern)
		if ok {
			envVars[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envVars, nil
}

// parseLine attempts to extract a key=value pair from a single line using regex.
func parseLine(line string, pattern *regexp.Regexp) (key, value string, ok bool) {
	matches := pattern.FindStringSubmatch(line)
	if len(matches) != 3 {
		return "", "", false
	}

	key = matches[1]
	value = cleanValue(matches[2])
	return key, value, true
}

// cleanValue trims and unquotes a value string.
func cleanValue(raw string) string {
	if len(raw) == 0 {
		return ""
	}

	raw = strings.TrimSpace(raw)
	if len(raw) >= 2 {
		if (raw[0] == '"' && raw[len(raw)-1] == '"') ||
			(raw[0] == '\'' && raw[len(raw)-1] == '\'') {
			return raw[1 : len(raw)-1]
		}
	}

	return raw
}
