package main

import (
	"context"
	"fmt"
	"fuzzer/modules/sql_injection"
	"fuzzer/utilities"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

type moduleFunc func(ctx context.Context, moduleConfig interface{}, apiUrl string, apiSchema string) error

type failedModule struct {
	moduleName string
	err        error
}

var availableModules = map[string]moduleFunc{
	"sql_injection": sql_injection.SQLInjectorModule,
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("No config file found. Existing\n")
	}
	fileName := args[1]
	config, err := ParseConfig(fileName)
	if err != nil {
		log.Fatalf("Unable to read config: %s", err)
	}

	apiSchema, err := readSchema(config.Runner.API.Schema)
	if err != nil {
		log.Fatalf("Unable to read schema: %s", err)
	}
	apiURL := generateAPIURL(config)
	log.Printf("Using API URL: %s", apiURL)

	var failedModules []failedModule
	for name := range config.Modules {
		moduleFunc, ok := availableModules[name]
		if !ok {
			log.Printf("Module not found: %s. Skipping\n", name)
			continue
		}
		err := runModule(name, config.Runner.ControlScript, moduleFunc, config.Modules[name], apiURL, apiSchema)
		if err != nil {
			failedModules = append(failedModules, failedModule{name, err})
		}
	}
	if len(failedModules) != 0 {
		for i := range failedModules {
			log.Printf("%s failed: %s\n", failedModules[i].moduleName, failedModules[i].err)
		}
		os.Exit(1)
	}
}

func readSchema(file string) (string, error) {
	schema, err := ioutil.ReadFile(file)
	return string(schema), err
}

func generateAPIURL(config *Config) string {
	url := url.URL{
		Scheme: config.Runner.API.HTTPScheme,
		Host:   fmt.Sprintf("%s:%d", config.Runner.API.Host, config.Runner.API.Port),
	}
	return url.String()
}

func runModule(moduleName, control string, module moduleFunc, moduleConfig interface{}, apiUrl, apiSchema string) error {
	defer func() {
		log.Printf("Module run complete. Shutting down current API/DB\n")
		err := utilities.StopAPI(control)
		if err != nil {
			log.Println(err)
		}
	}()
	log.Printf("Running module %s\nStarting API/DB for run\n", moduleName)
	if err := utilities.StartAPI(control); err != nil {
		log.Printf("unable to run control script to start API: %s", err)
		return err
	}
	// This can in theory be used to set global timeout limits
	ctx := context.Background()
	err := module(ctx,
		moduleConfig,
		apiUrl,
		apiSchema,
	)
	if err != nil {
		log.Printf("Received error from runner: %s", err)
	}
	return err
}
