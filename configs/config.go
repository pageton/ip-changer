package configs

import (
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
	ChangeIPOnRequest bool `yaml:"change_ip_on_request"`
}

func LoadConfig(path string) Config {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return config
}
