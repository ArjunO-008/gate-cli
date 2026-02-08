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

	_, err = os.Stat(gateConfigDirPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to check gate config directory: %w", err)
		}

		err = os.MkdirAll(gateConfigDirPath, 0700)
		if err != nil {
			return fmt.Errorf("failed to create gate config directory: %w", err)
		}

	}

	sampleConfigFilePath := filepath.Join(gateConfigDirPath, "sampleConfigFile.json")
	sampleConfigFile, err := os.OpenFile(sampleConfigFilePath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)

	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return fmt.Errorf("failed to create config file: %w", err)
	}

	defer sampleConfigFile.Close()

	config := models.Config{
		Version: "0.1.0",
		Name:    "sampleconfig",
		Settings: models.Settings{
			WorkingDirectory: ".",
		},
		Steps: []models.Step{
			{
				Type:    "executable",
				Command: "npm",
				Args:    "install",
				Dir:     "project",
			},
			{
				Type:    "shellexecutable",
				Command: "echo Build completed successfully",
			},
		},
		Extras: models.Extras{
			Git: models.GitExtras{
				Enabled: false,
				Config:  "gitConfig",
			},
		},
	}

	encoder := json.NewEncoder(sampleConfigFile)
	encoder.SetIndent("", " ")
	err = encoder.Encode(config)

	if err != nil {
		return fmt.Errorf("failed to write default config: %w", err)
	}

	return nil
}
