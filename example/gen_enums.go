// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/lggomez/go-enum/gen"
)

func main() {
	// Generate enums on an "enum" subpackage
	// To generate them on the current package, just use the current directory path (".")
	gen.GenerateEnumTypes(
		fmt.Sprintf(".%venum", string(os.PathSeparator)),
		"github.com/lggomez/go-enum/example/enum",
		false,
		gen.StringEnumDefinition{
			Name:   "Ghost",
			Values: []string{"Blinky", "Pinky", "Inky", "Clyde"},
		},
		gen.StringEnumDefinition{
			Name:   "Thing",
			Values: []string{"Foo", "Bar", "Baz", "Quux"},
		},
	)
}
