package cmd

import (
	"encoding/json"
	"fmt"
	"gate/models"
	"os"
	"path/filepath"
)

func InitCommandHandler() error {
	osConfigPath, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config directory: %w", err)
	}

	gateConfigDirPath := filepath.Join(osConfigPath, "gate-cli")

	err = os.MkdirAll(gateConfigDirPath, 0700)
	if err != nil {
		return fmt.Errorf("failed to create gate config directory: %w", err)
	}

	indexFilePath := filepath.Join(gateConfigDirPath, ".index.json")
	sampleConfigFilePath := filepath.Join(gateConfigDirPath, "sampleConfigFile.json")

	indexExists := true
	_, err = os.Stat(indexFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to check index file %q: %w", indexFilePath, err)
		}
		indexExists = false
	}

	if !indexExists {
		emptyIndex := map[string]string{}
		data, err := json.MarshalIndent(emptyIndex, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to serialize empty index data: %w", err)
		}

		err = os.WriteFile(indexFilePath, data, 0644)
		if err != nil {
			return fmt.Errorf("failed to create index file %q: %w", indexFilePath, err)
		}
	}

	_, err = os.Stat(sampleConfigFilePath)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check sample config file: %w", err)
	}

	config := models.Config{
		Version: "0.1.0",
		Name:    "sampleconfig",
		Settings: models.Settings{
			WorkingDirectory: ".",
		},
		Steps: []models.Step{
			{
				Type:    "executable",
				Command: "gate",
				Args:    "version",
				Dir:     ".",
			},
			{
				Type:    "shellexecutable",
				Command: "echo gate-cli runs successfully",
			},
		},
		
	}

	if config.Name == "" {
		return fmt.Errorf("sample config has empty name")
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize config: %w", err)
	}

	err = os.WriteFile(sampleConfigFilePath, data, 0600)
	if err != nil {
		return fmt.Errorf("failed to write sample config file: %w", err)
	}

	indexData, err := os.ReadFile(indexFilePath)
	if err != nil {
		return fmt.Errorf("failed to read index file: %w", err)
	}

	var index map[string]string
	err = json.Unmarshal(indexData, &index)
	if err != nil {
		return fmt.Errorf("failed to parse index file: %w", err)
	}

	index[config.Name] = sampleConfigFilePath

	updatedIndexData, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize updated index: %w", err)
	}

	err = os.WriteFile(indexFilePath, updatedIndexData, 0644)
	if err != nil {
		return fmt.Errorf("failed to update index file: %w", err)
	}

	return nil
}
