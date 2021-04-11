package sql_injection

import (
	"context"
	"fmt"
	"fuzzer/utilities"
	"io/ioutil"
	"log"
	"strconv"

	"gopkg.in/yaml.v2"
)

const dockerName = "fuzzer-modules-sql_injection"

type sqlInjectionConfig struct {
	APISchema string `yaml:"api_schema"`
}

func SQLInjectorModule(ctx context.Context, moduleConfig interface{}, apiPort, dbPort int, dbName, dbUsername, dbPassword string) error {
	// Kinda jank but works..
	data, _ := yaml.Marshal(moduleConfig)
	config := &sqlInjectionConfig{}
	yaml.Unmarshal(data, config)
	schemaData, err := ioutil.ReadFile(config.APISchema)
	if err != nil {
		return fmt.Errorf("unable to read schema file: %w", err)
	}
	env := map[string]string{
		"APIPort":    strconv.Itoa(apiPort),
		"DBPort":     strconv.Itoa(dbPort),
		"DBName":     dbName,
		"DBUsername": dbUsername,
		"DBPassword": dbPassword,
		"APISchema":  string(schemaData),
	}
	log.Printf("Starting SQL Injection Module Container\n")
	return utilities.RunImage(dockerName, env)
}
