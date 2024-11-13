package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DBConnectionString string `yaml:"db_connection_string"`
	ApiKey             string `yaml:"api_key"`
}

type Configs struct {
	Dev  Config `yaml:"dev"`
	Test Config `yaml:"test"`
	Prod Config `yaml:"prod"`
}

func LoadConfig(env string) (Config, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	// Parse the YAML data into a Configs struct
	var configs Configs
	if err := yaml.Unmarshal(data, &configs); err != nil {
		return Config{}, fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	// Return the config for the specified environment
	switch env {
	case "dev":
		return configs.Dev, nil
	case "test":
		return configs.Test, nil
	case "prod":
		return configs.Prod, nil
	default:
		return Config{}, fmt.Errorf("unknown environment: %s", env)
	}
}
