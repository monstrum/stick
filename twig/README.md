Twig
====

[![CircleCI](https://circleci.com/gh/monstrum/stick/tree/main.svg?style=shield)](https://circleci.com/gh/monstrum/stick/tree/main)
[![GoDoc](https://godoc.org/github.com/monstrum/stick/twig?status.svg)](https://godoc.org/github.com/monstrum/stick/twig)

Provides [Twig-compatibility](http://twig.sensiolabs.org/) for the stick
templating engine.


Overview
--------

This is the Twig compatibility subpackage for Stick.

### Current status

##### In development

Package
[`github.com/monstrum/stick/twig`](https://github.com/monstrum/stick/tree/master/twig)
contains extensions to provide the most Twig-like experience for
template writers. It aims to feature the same functions, filters, etc.
to be closely Twig-compatible.

Package
[`github.com/monstrum/stick`](https://github.com/monstrum/stick)
is a Twig template parser and executor. It provides the core
functionality and offers many of the same extension points as Twig like
functions, filters, node visitors, etc.


Installation
------------

The `twig` package is intended to be used as a library. The recommended
way to install the library is using `go get`.

```bash
go get -u github.com/monstrum/stick/twig
```


Usage
-----

Execute a simple Twig template.

```go
package main

import (
	"log"
	"os"
	
	"github.com/monstrum/stick"
	"github.com/monstrum/stick/twig"
)

func main() {
    env := twig.New(nil)
	if err := env.Execute("Hello, {{ name }}!", os.Stdout, map[string]stick.Value{"name": "Tyler"}); err != nil {
		log.Fatal(err)
	}
}
```

See [godoc for more information](https://pkg.go.dev/github.com/monstrum/stick/twig).

