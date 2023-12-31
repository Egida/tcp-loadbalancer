package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Client struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type ConfigFile struct {
	Config struct {
		Servers []Server `yaml:"servers"`
		Clients []Client `yaml:"clients"`
	} `yaml:"config"`
}

var cfg ConfigFile
var path string = "sample.yml"

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(fmt.Sprintf("%s/%s", currentDir, path))
	if err != nil {
		panic(err)
	}

	// Parse YAML data into the Config struct
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
}

func GetConfig() ConfigFile {
	return cfg
}
