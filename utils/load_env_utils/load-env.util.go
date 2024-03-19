package load_env_utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func LoadEnv() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// Absoulte Path for the env file from current file
	envFilePath := filepath.Join(basepath, "../../.env")

	// Read the environment file content
	fileContent, err := os.ReadFile(envFilePath)
	if err != nil {
		log.Fatal("Error Loading .env File", err)
	}

	// Split the content by lines after Removing trailing newline character
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	// Parse and set environment variables
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Invalid line format in environment file: %s\n", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
}
