package configs

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tor struct {
		ControlPort string `yaml:"control_port"`
	} `yaml:"tor"`
	Scheduler struct {
		Interval time.Duration `yaml:"interval"`
	} `yaml:"scheduler"`
}

func LoadConfig(path string) Config {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
	return config
}
