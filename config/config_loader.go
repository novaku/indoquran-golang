package config

import (
	"fmt"
	"os"

	"indoquran-golang/config/static"

	"gopkg.in/yaml.v2"
)

// LoadConfig : load the configuration file based on the ENV variables
func LoadConfig() *Config {
	var cfg Config
	readFile(&cfg)

	return &cfg
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	envVar := os.Getenv(static.EnvironmentVariableName)
	if envVar == "" {
		envVar = "development" // set default environment to development
	}

	filename := fmt.Sprintf("config/yaml/%s.yml", envVar)
	f, err := os.Open(filename)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
