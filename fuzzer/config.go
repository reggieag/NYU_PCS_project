package main

import (
	"fuzzer/pkg/runner"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Runner struct {
		ControlScript string `yaml:"control"`
		API           struct {
			Schema     string `yaml:"schema"`
			Host       string `yaml:"host"`
			Port       int    `yaml:"port"`
			HTTPScheme string `yaml:"http_scheme"`
			Security   struct {
				ClientsFile string `yaml:"clients_file"`
			}
		}
	}
	Modules       []map[string]interface{} `yaml:"modules"`
	ModulesParsed []runner.ModuleConfig
}

func parseConfig(file string) (*Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err2 := yaml.Unmarshal(data, config); err2 != nil {
		return nil, err2
	}
	config.ModulesParsed = make([]runner.ModuleConfig, 0, len(config.Modules))
	for i := range config.Modules {
		for key, data := range config.Modules[i] {
			config.ModulesParsed = append(config.ModulesParsed, runner.ModuleConfig{Name: key, Data: data})
			break
		}
	}
	return config, nil
}

func readClientFile(file string) (string, error) {
	if file == "" {
		return "", nil
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
