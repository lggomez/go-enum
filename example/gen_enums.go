// +build ignore

package main

import (
	"github.com/lggomez/go-enum/gen"
)

func main() {
	gen.GenerateEnumTypes("github.com/lggomez/go-enum/example", gen.StringEnumDefinition{
		Name:   "Cosas",
		Values: []string{"foo", "bar", "baz"},
	})
}
