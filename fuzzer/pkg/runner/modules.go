package runner

import "context"

type ModuleConfig struct {
	Name string
	Data interface{}
}

type ModuleErrors struct {
	Module ModuleConfig
	Error  error
}

type moduleFunc func(ctx context.Context, moduleConfig interface{}, apiUrl, apiSchema, clients string) error
