package sql_injection

import (
	"context"
	"fuzzer/utilities"
	"log"
	"strconv"

	"gopkg.in/yaml.v2"
)

const dockerName = "fuzzer-modules-sql_injection"

type sqlInjectionConfig struct {
	Exhaustive bool   `yaml:"exhaustive"`
	ForceHTTP  bool   `yaml:"force_http"`
	LogLevel   string `yaml:"log_level"`
	Database struct {
		Name     string
		Username string
		Password string
		Port     int
		Host     string
	}
}

func SQLInjectorModule(ctx context.Context, moduleConfig interface{}, apiUrl, apiSchema, clients string) error {
	// Kinda jank but works..
	data, _ := yaml.Marshal(moduleConfig)
	config := &sqlInjectionConfig{}
	yaml.Unmarshal(data, config)
	env := map[string]string{
		"API_URL":     apiUrl,
		"DB_PORT":     strconv.Itoa(config.Database.Port),
		"DB_NAME":     config.Database.Name,
		"DB_USERNAME": config.Database.Username,
		"DB_PASSWORD": config.Database.Password,
		"DB_HOST":     config.Database.Host,
		"API_SCHEMA":  apiSchema,
		"API_CLIENTS": clients,
		"EXHAUSTIVE":  strconv.FormatBool(config.Exhaustive),
		"FORCE_HTTP":  strconv.FormatBool(config.ForceHTTP),
		"LOG_LEVEL":   config.LogLevel,
	}
	log.Printf("Starting SQL Injection Module Container\n")
	return utilities.RunImage(dockerName, env)
}
