package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Runner struct {
		ComposeFile string `yaml:"compose_file"`
		Database    struct {
			Name     string
			Username string
			Password string
			Port     int
		}
		API struct {
			Port int
		}
	}
	Modules map[string]interface{}
}

func ParseConfig(file string) (*Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err2 := yaml.Unmarshal(data, config); err2 != nil {
		return nil, err2
	}
	return config, nil
}
