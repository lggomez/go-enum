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
			PackageDirectoryPath:  fmt.Sprintf(".%venum", string(os.PathSeparator)),
			PackageImportPath:     "github.com/lggomez/go-enum/example/enum",
			ValueIdentifierCasing: generator.CamelCase,
			OmitGeneratedNotice:   false,
			OmitTests:             false,
			OmitNameSanitization:  false,
			OmitSourceFormatting:  false,
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
			PackageDirectoryPath:  fmt.Sprintf(".%venum", string(os.PathSeparator)),
			PackageImportPath:     "github.com/lggomez/go-enum/example/enum",
			ValueIdentifierCasing: generator.UpperCase,
			OmitGeneratedNotice:   false,
			OmitTests:             false,
			OmitNameSanitization:  false,
			OmitSourceFormatting:  false,
		},
		generator.StringEnumDefinition{
			Name:   "CountriesISO3166-1",
			Values: []string{"Ca", "Uy", "Us", "Ar"},
		},
	)
}
