package models

type Config struct {
	Version  string   `json:"version"`
	Name     string   `json:"name"`
	Settings Settings `json:"settings"`
	Steps    []Step   `json:"steps"`
}

type Settings struct {
	WorkingDirectory string `json:"workingDirectory"`
}

type Step struct {
	Type    string `json:"type"`
	Command string `json:"command"`
	Args    string `json:"args,omitempty"`
	Dir     string `json:"dir,omitempty"`
}
