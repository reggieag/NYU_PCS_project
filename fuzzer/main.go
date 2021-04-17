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

type moduleFunc func(ctx context.Context, moduleConfig interface{}, apiUrl, apiSchema, clients string) error

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
		log.Fatalf("No config file found. Exiting\n")
	}
	fileName := args[1]
	config, err := parseConfig(fileName)
	if err != nil {
		log.Fatalf("Unable to read config: %s", err)
	}

	apiSchema, err := readSchema(config.Runner.API.Schema)
	if err != nil {
		log.Fatalf("Unable to read schema: %s", err)
	}
	apiURL := generateAPIURL(config)
	log.Printf("Using API URL: %s", apiURL)

	log.Printf("Reading clients file")
	clients, err := readClientFile(config.Runner.API.Security.ClientsFile)
	if err != nil {
		log.Fatalf("Unable to read clients: %s", err)
	}

	var failedModules []failedModule
	for _, module := range config.ModulesParsed {
		moduleFunc, ok := availableModules[module.Name]
		if !ok {
			log.Printf("Module not found: %s. Skipping\n", module.Name)
			continue
		}
		err := runModule(module.Name, config.Runner.ControlScript, moduleFunc, module.Data, apiURL, apiSchema, clients)
		if err != nil {
			failedModules = append(failedModules, failedModule{module.Name, err})
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

func runModule(moduleName, control string, module moduleFunc, moduleConfig interface{}, apiUrl, apiSchema, clients string) error {
	defer func() {
		log.Println("Module run complete")
		log.Println("Calling control script with argument 'stop'")
		err := utilities.StopAPI(control)
		if err != nil {
			log.Println(err)
		}
	}()
	log.Printf("Running module %s", moduleName)
	log.Println("Calling control script with argument 'start' for run:")
	if err := utilities.StartAPI(control); err != nil {
		log.Printf("unable to run control script to start API: %s", err)
		return err
	}
	log.Println("Control script finished")
	// This can in theory be used to set global timeout limits
	ctx := context.Background()
	err := module(ctx,
		moduleConfig,
		apiUrl,
		apiSchema,
		clients,
	)
	if err != nil {
		log.Printf("Received error from runner: %s", err)
	}
	return err
}
