package settings

import (
	_ "embed"
	"github.com/go-yaml/yaml"
)

//go:embed settings.yml
var settingsFile []byte

type DatabaseConfig struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Settings struct {
	Port string         `yaml:"port"`
	DB   DatabaseConfig `yaml:"database"`
}

func New() (*Settings, error) {
	var s Settings

	err := yaml.Unmarshal(settingsFile, &s) // ---> load in settingsFile
	if err != nil {
		return nil, err
	}
	return &s, nil
}
