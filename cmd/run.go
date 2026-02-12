package cmd

import (
	"encoding/json"
	"fmt"
	"gate/models"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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
	err = executeConfig(configFilePath)
	if err != nil {
		return err
	}

	return nil
}

func executeConfig(userConfigFile string) error {

	data, err := os.ReadFile(userConfigFile)
	if err != nil {
		return fmt.Errorf("failed to read config file %q: %w", userConfigFile, err)
	}

	var config models.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("invalid JSON in config file %q: %w", userConfigFile, err)
	}

	if len(config.Steps) == 0 {
		return fmt.Errorf("no steps found for config %q", config.Name)
	}

	workingDir := config.Settings.WorkingDirectory
	if workingDir == "" {
		workingDir = "."
	}

	for _, step := range config.Steps {

		switch step.Type {
		case "executable":
			if err := runExecutable(step, workingDir); err != nil {
				return fmt.Errorf("step %q failed: %w", step, err)
			}

		case "shellexecutable":
			if err := runShellexecutable(step, workingDir); err != nil {
				return fmt.Errorf("step %q failed: %w", step, err)
			}

		default:
			return fmt.Errorf("invalid step type %q", step.Type)

		}

	}

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
		return "", fmt.Errorf("failed to stat .index.json: %w", err)
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

func runExecutable(step models.Step, workingDir string) error {
	args := strings.Fields(step.Args)
	exeCmd := exec.Command(step.Command, args...)

	if step.Dir != "" {
		exeCmd.Dir = filepath.Join(workingDir, step.Dir)
	} else {
		exeCmd.Dir = workingDir
	}

	exeCmd.Stdout = os.Stdout
	exeCmd.Stderr = os.Stderr
	exeCmd.Stdin = os.Stdin
	fmt.Printf("Running: %s %s\n", step.Command, step.Args)
	err := exeCmd.Run()
	if err != nil {
		return fmt.Errorf("command %q failed: %w", step.Command, err)
	}

	return nil
}
func runShellexecutable(step models.Step, workingDir string) error {
	var shellCmd *exec.Cmd

	if runtime.GOOS == "windows" {
		shellCmd = exec.Command("cmd", "/C", step.Command)
	} else {
		shellCmd = exec.Command("sh", "-c", step.Command)
	}

	if step.Dir != "" {
		shellCmd.Dir = filepath.Join(workingDir, step.Dir)
	} else {
		shellCmd.Dir = workingDir
	}

	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr
	shellCmd.Stdin = os.Stdin

	err := shellCmd.Run()
	if err != nil {
		return fmt.Errorf("shell command %q failed: %w", step.Command, err)
	}

	return nil
}
