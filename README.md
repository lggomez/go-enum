![Go](https://github.com/lggomez/go-enum/workflows/Go/badge.svg?branch=master)
[![GoDoc](https://godoc.org/github.com/lggomez/go-enum?status.svg)](https://pkg.go.dev/github.com/lggomez/go-enum?tab=doc)
[![Release](https://img.shields.io/github/release/lggomez/go-enum.svg?style=flat-square)](https://github.com/lggomez/go-enum/releases)

# go-enum -  Typesafe enum generation in golang

This package provides a scaffolding mechanism to generate enum types based on string values, in contrast with the _iota_ based integer constants enum pattern commonly used in Go

Use this package if you wish to have type support for:
* Traversable enumerations
* Enum fields compatible with several (un)marshal interfaces:
    * Stringer
    * json.Marshaler, json.Unmarshaler
    * text.Marshaler, text.Unmarshaler
    * json.Marshaler, json.Unmarshaler
    * gob.GobEncoder, gob.GobDecoder
    * driver.Valuer, sql.Scanner
    * bson.Marshaler, bson.Unmarshaler (from [go.mongodb.org/mongo-driver/bson](https://godoc.org/go.mongodb.org/mongo-driver/bson) package)
* Ability to perform type-safe comparisons at runtime against strings and instances of the same enum type

For more information on the enumeration type API, see the [example](https://pkg.go.dev/github.com/lggomez/go-enum@v0.4.0/example/enum?tab=doc#SpecialThing):

## Usage

This generator uses the `go generate` command, and as such, its pattern must be used in order to generate code

As shown in the example package, 2 files are needed:

#### example.go
```go
package example

//go:generate go run gen_enums.go
```

#### gen_enums.go
```go
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/lggomez/go-enum/generator"
)

// This example generates the 'Ghost', 'Thing' and 'CountriesIso31661' enums inside of an 'enum' subpackage
// To generate them on the current package instead, just use the current directory path (".")
func main() {
	generator.GenerateEnumTypes(
		generator.Options{
			PackageDirectoryPath: fmt.Sprintf(".%venum", string(os.PathSeparator)),
			PackageImportPath:    "github.com/lggomez/go-enum/example/enum",
			OmitGeneratedNotice:  false,
			OmitTests:            false,
			OmitNameSanitization: false,
			OmitSourceFormatting: false,
		},
		generator.StringEnumDefinition{
			Name:   "Ghost",
			Values: []string{"Blinky", "Pinky", "Inky", "Clyde"},
		},
		generator.StringEnumDefinition{
			Name:   "SpecialThing",
			Values: []string{"Foo", "Bar", "Baz", "Quux"},
		},
	)
	generator.GenerateEnumTypes(
		generator.Options{
			PackageDirectoryPath: fmt.Sprintf(".%venum", string(os.PathSeparator)),
			PackageImportPath:    "github.com/lggomez/go-enum/example/enum",
			ValueIdentifierCasing: generator.UpperCase,
			OmitGeneratedNotice:  false,
			OmitTests:            false,
			OmitNameSanitization: false,
			OmitSourceFormatting: false,
		},
		generator.StringEnumDefinition{
			Name:   "CountriesISO3166-1",
			Values: []string{"Ca", "Uy", "Us", "Ar"},
		},
	)
}

```

With both files present in the same directory/package, execute the `go generate` command on that directory

Based on this example, it will generate an `enum` subpackage with the `Ghost` and `Thing` enum types, along with its iterators `EnumGhost` and `EnumThing`
