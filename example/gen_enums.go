// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/lggomez/go-enum/generator"
)

func main() {
	// This example generates the 'Ghost' and 'Thing' enums on an 'enum' subpackage
	// To generate them on the current package, just use the current directory path (".")
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
			Name:   "Thing",
			Values: []string{"Foo", "Bar", "Baz", "Quux"},
		},
	)
}
