package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ConfigCommandHandler(arguments []string) error {

	if len(arguments) == 0 {
		return fmt.Errorf("no config subcommand provided")
	}

	firstArgument := arguments[0]

	switch firstArgument {
	case "path":
		showConfigPath()
	case "list":
		showConfiFileList()
	default:
		return fmt.Errorf("unknown config subcommand: %s", firstArgument)
	}

	return nil
}

func showConfigPath() error {
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

	fmt.Println(configFilePath)
	return nil
}

func showConfiFileList() error {
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

	configFiles, err := os.ReadDir(configFilePath)
	if err != nil {
		return err
	}

	for _, file := range configFiles {
		name := file.Name()

		if strings.HasPrefix(name, ".") {
			continue
		}
		fmt.Println(name)
	}

	return nil

}
