package runner

import (
	"context"
	"fmt"
)

type StartSetup func(ctx context.Context, module ModuleConfig) error
type StopTeardown func(ctx context.Context, module ModuleConfig, runError error)

type Runner struct {
	schema      string
	apiURL      string
	clientsList string
	modules     []ModuleConfig
}

func NewRunner(schema, apiURL, clientsList string, modules []ModuleConfig) *Runner {
	return &Runner{
		schema:      schema,
		apiURL:      apiURL,
		clientsList: clientsList,
		modules:     modules,
	}
}

func (r *Runner) Execute(ctx context.Context, start StartSetup, stop StopTeardown) []ModuleErrors {
	errors := make([]ModuleErrors, 0, len(r.modules))
	for _, module := range r.modules {
		moduleFunc, ok := availableModules[module.Name]
		if !ok {
			errors = append(errors, ModuleErrors{
				Module: module,
				Error:  fmt.Errorf("error runner module %s: %w", module.Name, ModuleNotFound),
			})
			continue
		}
		err := r.executeModule(ctx, start, stop, moduleFunc, module)
		errors = append(errors, ModuleErrors{
			Module: module,
			Error:  err,
		})
	}
	return errors
}

func (r *Runner) executeModule(ctx context.Context, start StartSetup, stop StopTeardown, module moduleFunc, moduleConfig ModuleConfig) error {
	var err error
	defer stop(ctx, moduleConfig, err)
	err = start(ctx, moduleConfig)
	if err != nil {
		return err
	}
	err = module(ctx, moduleConfig.Data, r.apiURL, r.schema, r.clientsList)
	return err
}
