package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func DeleteCommandHandler(argument []string) error {
	if len(argument) == 0 {
		return fmt.Errorf("no file to delete found")
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

	for _, name := range argument {
		path, ok := index[name]
		if !ok {
			return fmt.Errorf("name %q not found in index.json", name)
		}
		if _, err := os.Stat(path); err == nil {
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("failed to delete file %q: %w", path, err)
			}
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("failed to check file %q: %w", path, err)
		}
		delete(index, name)
		updatedData, err := json.MarshalIndent(index, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal index.json: %w", err)
		}

		if err := os.WriteFile(indexPath, updatedData, 0644); err != nil {
			return fmt.Errorf("failed to write index.json: %w", err)
		}

	}

	return nil
}
