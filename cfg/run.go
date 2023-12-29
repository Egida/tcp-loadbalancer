package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigFile struct {
	Config struct {
		Servers []struct {
			Name string `yaml:"name"`
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"servers"`
		Clients []struct {
			Name string `yaml:"name"`
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"clients"`
	} `yaml:"config"`
}

var config Config
var path string = "sample.yaml"

func init() {

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Parse YAML data into the Config struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
	fmt.Println(config)
}
