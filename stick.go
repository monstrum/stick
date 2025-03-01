package stick // import "github.com/monstrum/stick"

import (
	"bytes"
	"context"
	"io"

	"github.com/monstrum/stick/parse"
)

// A Template represents a named template and its contents.
type Template interface {
	// Name returns the name of this Template.
	Name() string

	// Contents returns an io.Reader for reading the Template contents.
	Contents() io.Reader
}

// A Func represents a user-defined function.
// Functions can be called anywhere expressions are allowed and
// take any number of arguments.
type Func func(ctx Context, args ...Value) Value

// A Filter is a user-defined filter.
// Filters receive a value and modify it in some way. Filters
// also accept parameters.
type Filter func(ctx Context, val Value, args ...Value) Value

// A Test represents a user-defined test.
// Tests are used to make some comparisons more expressive. Tests
// also accept arguments and can consist of two words.
type Test func(ctx Context, val Value, args ...Value) bool

// Env represents a configured Stick environment.
type Env struct {
	Loader    Loader              // Template loader.
	Functions map[string]Func     // User-defined functions.
	Filters   map[string]Filter   // User-defined filters.
	Tests     map[string]Test     // User-defined tests.
	Visitors  []parse.NodeVisitor // User-defined node visitors.
	Globals   map[string]Value    // Global values available to all templates.
}

// An Extension is used to group related functions, filters, visitors, etc.
type Extension interface {
	// Init is the entry-point for an extension to modify the Env.
	Init(*Env) error
}

// ContextMetadata contains additional, unstructured runtime attributes about
// the template being executed.
type ContextMetadata interface {
	All() map[string]string         // Returns a map of all attributes and values.
	Set(name, val string)           // Set a metadata attribute on the context.
	Get(name string) (string, bool) // Get a metadata attribute on the context.

	noexport() // Prevent other packages from satisfying this interface.
}

// ContextScope provides an interface with the currently executing template's
// scope.
type ContextScope interface {
	All() map[string]Value    // Returns a map of all values defined in the scope.
	Get(string) (Value, bool) // Get a value defined in the scope.
	Set(string, Value)        // Set a value in the scope.

	noexport() // Prevent other packages from satisfying this interface.
}

// A Context represents the execution context of a template.
//
// The Context is passed to all user-defined functions, filters, tests,
// and node visitors. It can be used to affect and inspect the local
// environment while a template is executing.
type Context interface {
	Name() string          // The name of the template being executed.
	Meta() ContextMetadata // Runtime metadata about the template.
	Scope() ContextScope   // All defined root-level names.
	Env() *Env
	Context() context.Context // The context.Context of the request.
	noexport()                // Prevent other packages from satisfying this interface.
}

// New creates an empty Env.
// If nil is passed as loader, a StringLoader is used.
func New(loader Loader) *Env {
	if loader == nil {
		loader = &StringLoader{}
	}
	return &Env{
		Loader:    loader,
		Functions: make(map[string]Func),
		Filters:   make(map[string]Filter),
		Tests:     make(map[string]Test),
		Visitors:  make([]parse.NodeVisitor, 0),
		Globals:   make(map[string]Value),
	}
}

// Register adds the given Extension to the Env.
func (env *Env) Register(e Extension) error {
	return e.Init(env)
}

// Execute parses and executes the given template.
func (env *Env) Execute(ctx context.Context, tpl string, out io.Writer, values map[string]Value) error {
	return execute(ctx, tpl, out, values, env)
}

// ExecuteSafe executes the template but does not output anything if an error occurs.
func (env *Env) ExecuteSafe(ctx context.Context, tpl string, out io.Writer, values map[string]Value) error {
	buf := &bytes.Buffer{}
	if err := env.Execute(ctx, tpl, buf, values); err != nil {
		return err
	}
	_, err := io.Copy(out, buf)
	return err
}

// Parse loads and parses the given template.
func (env *Env) Parse(name string) (*parse.Tree, error) {
	return env.load(name)
}
