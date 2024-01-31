[![codecov](https://codecov.io/gh/gofika/jsonutil/branch/main/graph/badge.svg)](https://codecov.io/gh/gofika/jsonutil)
[![Build Status](https://github.com/gofika/jsonutil/workflows/build/badge.svg)](https://github.com/gofika/jsonutil)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofika/jsonutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofika/jsonutil)](https://goreportcard.com/report/github.com/gofika/jsonutil)
[![Licenses](https://img.shields.io/github/license/gofika/jsonutil)](LICENSE)

# jsonutil

golang json utils for common use

## Basic Usage

### Installation

To get the package, execute:

```bash
go get github.com/gofika/jsonutil
```

### Example

```go
package main

import (
	"fmt"
	"github.com/gofika/jsonutil"
)

type Foo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	name := "foo.json"
	// write struct to file
	err := jsonutil.WriteFile(name, foo)
	if err != nil {
		fmt.Printf("WriteFile failed. err: %s\n", err.Error())
		return
	}
	// read struct from file
	bar, err = jsonutil.ReadFile[Foo](name)
	if err != nil {
		fmt.Printf("ReadFile failed. err: %s\n", err.Error())
		return
	}
	fmt.Printf("bar.Name: %s\n", bar.Name)
	fmt.Printf("bar.Value: %d\n", bar.Value)
}
```