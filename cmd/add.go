package cmd

import (
	"encoding/json"
	"fmt"
	"gate/models"
	"os"
	"path/filepath"
)

func AddCommandHandler(filePaths []string) error {
	if len(filePaths) == 0 {
		return fmt.Errorf("no add subcommand provided")
	}

	const gateDirName = "gate-cli"

	osConfigPath, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config directory: %w", err)
	}

	configFilePath := filepath.Join(osConfigPath, gateDirName)

	info, err := os.Stat(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config directory not found, run `gate init` first")
		}

		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("%s exists but is not a directory", configFilePath)
	}

	err = addFileToConfigDir(filePaths, configFilePath)
	if err != nil {
		return err
	}

	return nil
}

func validateJSON(userJsonFile string) (string, error) {

	if filepath.Ext(userJsonFile) != ".json" {
		return "", fmt.Errorf("file %s is not a .json file", userJsonFile)
	}

	fileInfo, err := os.Stat(userJsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file %s does not exist", userJsonFile)
		}

		return "", fmt.Errorf("cannot access file %s: %w", userJsonFile, err)

	}

	if fileInfo.IsDir() {
		return "", fmt.Errorf("%s is a directory, not a file", userJsonFile)
	}

	jsonContent, err := os.ReadFile(userJsonFile)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", userJsonFile, err)
	}

	var fileContent models.Config

	err = json.Unmarshal(jsonContent, &fileContent)
	if err != nil {
		return "", fmt.Errorf("file %s is not valid JSON: %w", userJsonFile, err)
	}

	if fileContent.Name == "" {
		return "", fmt.Errorf("file %s does not have a valid 'name' field", userJsonFile)
	}

	return fileContent.Name, nil
}

func addFileToConfigDir(files []string, configFilePath string) error {

	indexPath := filepath.Join(configFilePath, ".index.json")

	var index map[string]string

	indexFileData, err := os.ReadFile(indexPath)
	if err != nil {
		if os.IsNotExist(err) {
			index = make(map[string]string)
		} else {
			return fmt.Errorf("failed to read index.json: %w", err)
		}
	} else {
		err = json.Unmarshal(indexFileData, &index)
		if err != nil {
			return fmt.Errorf("failed to parse index.json: %w", err)
		}
	}

	for _, file := range files {
		name, err := validateJSON(file)
		if err != nil {
			return err
		}

		fileName := filepath.Base(file)
		destPath := filepath.Join(configFilePath, fileName)

		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", file, err)
		}

		err = os.WriteFile(destPath, content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file to config dir: %w", err)
		}

		absPath, err := filepath.Abs(destPath)
		if err != nil {
			return fmt.Errorf("failed to resolve absolute path: %w", err)
		}

		index[name] = absPath
		fmt.Printf("Success fully added %s as %s\n", file, name)

	}

	updatedIndexData, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal index: %w", err)
	}

	err = os.WriteFile(indexPath, updatedIndexData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write index.json: %w", err)
	}

	return nil
}
