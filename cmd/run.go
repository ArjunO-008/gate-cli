package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func RunCommandHandler(arguments []string) error {
	if len(arguments) == 0 {
		return fmt.Errorf("no add subcommand provided")
	}
	fileIdentifier := arguments[0]
	configFilePath, err := readtIndexFile(fileIdentifier)
	if err != nil {
		return err
	}
	fmt.Println(configFilePath)

	return nil
}

func executeConfig(userConfigFile string) error {
	return nil
}

func readtIndexFile(fileIdentifier string) (string, error) {
	const gateDirName = "gate-cli"

	osConfigPath, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user config directory: %w", err)
	}

	configFilePath := filepath.Join(osConfigPath, gateDirName)

	info, err := os.Stat(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("config directory not found, run `gate init` first")
		}

		return "", err
	}

	if !info.IsDir() {
		return "", fmt.Errorf("%s exists but is not a directory", configFilePath)
	}

	indexPath := filepath.Join(configFilePath, ".index.json")
	_, err = os.Stat(indexPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf(".index.json does not exist, run gate init")
		}
		return "", nil
	}

	userConfigIdentifier, err := os.ReadFile(indexPath)
	if err != nil {
		return "", fmt.Errorf("failed to read .index.json: %w", err)
	}
	index := make(map[string]string)
	err = json.Unmarshal(userConfigIdentifier, &index)
	if err != nil {
		return "", fmt.Errorf("invalid .index.json format: %w", err)
	}

	path, ok := index[fileIdentifier]
	if !ok {
		return "", fmt.Errorf("no entry found for name: %s", fileIdentifier)
	}

	return path, nil

}
