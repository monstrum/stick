package factory

import (
	"github.com/monstrum/stick"
	"github.com/monstrum/stick/parse"
	"github.com/monstrum/stick/twig/filter"
)

type EnvOption func(*stick.Env)

type AppendFilterFn func(filters map[string]stick.Filter)

type AppendFunctionFn func(filters map[string]stick.Func)

func New(
	options ...EnvOption,
) *stick.Env {
	env := &stick.Env{}
	for _, option := range options {
		option(env)
	}
	return env
}

func WithLoader(loader stick.Loader) EnvOption {
	return func(env *stick.Env) {
		env.Loader = loader
	}
}

func WithFunctions(functions map[string]stick.Func, additionalFunctions ...AppendFunctionFn) EnvOption {
	for _, fn := range additionalFunctions {
		fn(functions)
	}
	return func(env *stick.Env) {
		env.Functions = functions
	}
}

func WithFilters(filters map[string]stick.Filter, additionalFilters ...AppendFilterFn) EnvOption {
	for _, fn := range additionalFilters {
		fn(filters)
	}
	return func(env *stick.Env) {
		env.Filters = filters
	}
}

func WithTests(tests map[string]stick.Test) EnvOption {
	return func(env *stick.Env) {
		env.Tests = tests
	}
}

func WithVisitors(visitors []parse.NodeVisitor) EnvOption {
	return func(env *stick.Env) {
		env.Visitors = visitors
	}
}

func WithDefaultFilters(additionalFilters ...AppendFilterFn) EnvOption {
	filters := filter.TwigFilters()
	for _, fn := range additionalFilters {
		fn(filters)
	}
	return func(env *stick.Env) {
		env.Filters = filters
	}
}
