package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func EditCommandHandler(fileName string) error {
	if fileName == "" {
		return fmt.Errorf("no filename found")
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

	path, ok := index[fileName]
	if !ok {
		return fmt.Errorf("name %q not found in index.json", fileName)
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", path, err)
	}

	return nil
}
