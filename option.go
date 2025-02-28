package stick

import (
	"context"
	"log/slog"

	"github.com/monstrum/stick/parse"
)

// Option is a configuration option for an Env.
type Option func(context context.Context, logger *slog.Logger, env *Env)

func WithOptions(ctx context.Context, logger *slog.Logger, opts ...Option) *Env {
	env := &Env{
		Loader:    &StringLoader{},
		Globals:   make(map[string]Value),
		Functions: make(map[string]Func),
		Filters:   make(map[string]Filter),
		Tests:     make(map[string]Test),
		Visitors:  make([]parse.NodeVisitor, 0),
	}
	for _, opt := range opts {
		opt(ctx, logger, env)
	}
	return env
}

// WithLoader sets the loader for the environment.
func WithLoader(loader Loader) Option {
	return func(ctx context.Context, logger *slog.Logger, env *Env) {
		env.Loader = loader
	}
}

func WithFunctions(
	funcNames []string,
	callback []Func,
) Option {
	return func(ctx context.Context, logger *slog.Logger, env *Env) {
		if len(funcNames) != len(callback) {
			logger.ErrorContext(ctx, "funcNames and callback must have the same length")
			return
		}
		for i := range funcNames {
			env.Functions[funcNames[i]] = callback[i]
		}
	}
}

func WithFilters(
	filterNames []string,
	filters []Filter,
) Option {
	return func(ctx context.Context, logger *slog.Logger, env *Env) {
		if len(filterNames) != len(filters) {
			logger.ErrorContext(ctx, "filterNames and filters must have the same length")
			return
		}
		for i := range filterNames {
			env.Filters[filterNames[i]] = filters[i]
		}
	}
}

func WithTests(
	testNames []string,
	tests []Test,
) Option {
	return func(ctx context.Context, logger *slog.Logger, env *Env) {
		if len(testNames) != len(tests) {
			logger.ErrorContext(ctx, "testNames and tests must have the same length")
			return
		}

		for i := range testNames {
			env.Tests[testNames[i]] = tests[i]
		}
	}
}

func WithGlobals(data ...interface{}) Option {
	return func(ctx context.Context, logger *slog.Logger, env *Env) {
		if len(data)%2 != 0 {
			logger.ErrorContext(ctx, "WithGlobals", "error", "data must be a multiple of 2")
			return
		}

		for i := 0; i < len(data); i += 2 {
			env.Globals[data[i].(string)] = Value(data[i+1])
		}
	}
}

func WithVisitors(
	visitors ...parse.NodeVisitor,
) Option {
	return func(ctx context.Context, logger *slog.Logger, env *Env) {
		env.Visitors = append(env.Visitors, visitors...)
	}
}

func WithKeys(callbacks ...string) []string {
	return callbacks
}

func Functions(functions ...Func) []Func {
	return functions
}

func Filters(filters ...Filter) []Filter {
	return filters
}
