package generator

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/stoewer/go-strcase"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/lggomez/go-enum/generator/internal"
	"github.com/lggomez/go-enum/generator/internal/templates"
)

// StringEnumDefinition is the basic [name:values] definition of an enumeration.
// As the name implies, this is for string enumerations only
type StringEnumDefinition struct {
	Name   string
	Values []string
}

// Options defined the settings to be passed to the generator in order to configure
// certain aspects of the code generation
type Options struct {
	// Filesystem path of the directory corresponding to the package to be used or created
	PackageDirectoryPath string
	// Import path of the package to be used or created. It must be a valid path according to the working module structure
	PackageImportPath string
	// Whether to omit the generated code header on files. Default value is false
	OmitGeneratedNotice bool
	// Whether to omit tests for generated code. Default value is false
	OmitTests bool
	// Whether to omit generated source code formatting, which also detects compilation errors. Default value is false
	OmitSourceFormatting bool
	// Whether to omit field name sanitization, which prevents invalid name qualifiers during code generation. Default value is false
	OmitNameSanitization bool
}

// canonicalStringEnum contains the full metadata required to execute the code templates and generate the specific implementations
type canonicalStringEnum struct {
	StructName          string
	StructNameLowerCase string

	IndexKeyName string
	Values       map[string]string

	TestCase testFieldMetadata

	ImportPath          string
	FileName            string
	Timestamp           string
	Package             string
	OmitGeneratedNotice bool
}

type testFieldMetadata struct {
	Name               string
	Key                string
	Value              string
	ValueScrambledCase string
	InvalidValue       string
	BinaryLen          int
	BSONLen            int
}

type fieldMetadata struct {
	Value    string
	TestCase testFieldMetadata
}

// GenerateEnumTypes scaffolds enum types for the given options and definitions
func GenerateEnumTypes(options Options, enums ...StringEnumDefinition) {
	if len(enums) == 0 {
		log.Panic("generator: enums are required")
	}

	// Get package name from import path
	// i.e: github.com/lggomez/go-enum/example -> example
	tokens := strings.Split(options.PackageImportPath, "/")
	packageName := tokens[len(tokens)-1]

	// Convert enum definitions into canonical definitions with full metadata for code generation
	canonicalEnums := processEnumerations(options.PackageImportPath, packageName, options, enums)

	if _, err := os.Stat(options.PackageDirectoryPath); os.IsNotExist(err) {
		if dirErr := os.Mkdir(options.PackageDirectoryPath, os.ModePerm); dirErr != nil {
			log.Panic("generator: could not create package - ", err.Error())
		}
	}

	// Traverse enums for code generation
	for i, canonicalEnum := range canonicalEnums {
		// Generate base enum struct and its codecs
		// This is a single time pass that must be done on the first iteration
		if i == 0 {
			if err := generateFileFromTemplate(options,
				canonicalEnum,
				templates.EnumTemplate,
				fmt.Sprintf("%s%senum.go", options.PackageDirectoryPath, string(os.PathSeparator))); err != nil {
				log.Panic(err.Error())
			}

			if err := generateFileFromTemplate(options,
				canonicalEnum,
				templates.EnumCodecsTemplate,
				fmt.Sprintf("%s%senum_codecs.go", options.PackageDirectoryPath, string(os.PathSeparator))); err != nil {
				log.Panic(err.Error())
			}
		}

		// Generate specific enum implementation file
		if err := generateFileFromTemplate(options,
			canonicalEnum,
			templates.EnumImplTemplate,
			fmt.Sprintf("%s%s%s.go", options.PackageDirectoryPath, string(os.PathSeparator), canonicalEnum.FileName)); err != nil {
			log.Panic(err.Error())
		}

		if !options.OmitTests {
			// Generate specific enum implementation test file
			if err := generateFileFromTemplate(options,
				canonicalEnum,
				templates.EnumImplTestTemplate,
				fmt.Sprintf("%s%s%s_test.go", options.PackageDirectoryPath, string(os.PathSeparator), canonicalEnum.FileName)); err != nil {
				log.Panic(err.Error())
			}

			// Generate codecs and base stringEnumValue test files
			// This is a single time pass that must be done on the first iteration,
			// and after the first specific enum is generated since it uses it for tests
			if i == 0 {
				if err := generateFileFromTemplate(options,
					canonicalEnum,
					templates.EnumTestTemplate,
					fmt.Sprintf("%s%senum_test.go", options.PackageDirectoryPath, string(os.PathSeparator))); err != nil {
					log.Panic(err.Error())
				}

				if err := generateFileFromTemplate(options,
					canonicalEnum,
					templates.EnumCodecsTestTemplate,
					fmt.Sprintf("%s%senum_codecs_test.go", options.PackageDirectoryPath, string(os.PathSeparator))); err != nil {
					log.Panic(err.Error())
				}
			}
		}
	}
}

func generateFileFromTemplate(options Options, canonicalEnum canonicalStringEnum, templateString, destinationPath string) error {
	// Generate code from template and save it to buffer
	src := &bytes.Buffer{}
	enumTemplate := template.Must(template.New(destinationPath).Parse(templateString))
	err := enumTemplate.Execute(src, canonicalEnum)
	if err != nil {
		log.Panicf("generator: could not create source from template (destination %s): %s", destinationPath, err.Error())
	}

	code := src.Bytes()

	// Run go fmt formatting into buffered source
	if !options.OmitSourceFormatting {
		code, err = internal.FormatSource(bytes.NewBuffer(code))
		if err != nil {
			log.Panicf("generator: could not format source (destination %s): %s", destinationPath, err.Error())
		}
	}

	// Write generated source to disk file
	err = internal.SaveFile(destinationPath, code)
	if err != nil {
		log.Panicf("generator: could not save source to file (destination %s): %s", destinationPath, err.Error())
	}

	return nil
}

func processEnumerations(importPath string, packageName string, options Options, enums []StringEnumDefinition) []canonicalStringEnum {
	canonicalEnums := make([]canonicalStringEnum, len(enums))

	for i, e := range enums {
		upperCamelName := strcase.UpperCamelCase(e.Name)
		lowerCamelName := strcase.LowerCamelCase(e.Name)

		ce := &canonicalStringEnum{
			Package:             packageName,
			ImportPath:          importPath,
			StructName:          upperCamelName,
			StructNameLowerCase: lowerCamelName,
			IndexKeyName:        lowerCamelName + "Key",
			Values:              map[string]string{},
			TestCase: testFieldMetadata{
				Key:          strcase.SnakeCase(e.Name),
				InvalidValue: "invalid_" + uuid.New().String(), // Set an unique random value to prevent collisions
			},
			FileName:            strcase.SnakeCase(e.Name),
			Timestamp:           time.Now().Format(time.RFC3339),
			OmitGeneratedNotice: options.OmitGeneratedNotice,
		}

		if len(enums) == 0 {
			log.Panic("generator: invalid zero length enum" + e.Name)
		}

		// traverse enum values and generate the field metadata for each value and its test
		for i, value := range e.Values {
			v := value
			if !options.OmitNameSanitization {
				// sanitize the struct name so it is a valid Go identifier
				v = sanitizeNameQualifier(v)
			}

			valueKey := strcase.UpperCamelCase(v)
			if i == 0 {
				// generate metadata of test case value from first value of the list
				ce.TestCase.Name = ce.StructName + valueKey
				ce.TestCase.Value = value
				ce.TestCase.ValueScrambledCase = scrambleCase(value)
				ce.TestCase.BinaryLen = len([]byte(ce.TestCase.Value))
				ce.TestCase.BSONLen = calculateBSONLen(ce.TestCase.Value)
			}

			ce.Values[valueKey] = value
		}

		canonicalEnums[i] = *ce
	}

	return canonicalEnums
}

func scrambleCase(value string) string {
	s := ""

	for _, c := range value {
		if unicode.IsUpper(c) {
			s = s + strings.ToLower(string(c))
		} else {
			s = s + strings.ToUpper(string(c))
		}
	}

	return s
}

func sanitizeNameQualifier(value string) string {
	camelRe := regexp.MustCompile(`[_. ]`)
	eraseRe := regexp.MustCompile(`[^a-zA-Z\d_. ]`)

	firstPass := camelRe.ReplaceAllString(value, "_")
	return eraseRe.ReplaceAllString(firstPass, "")
}

func calculateBSONLen(value string) int {
	v, _ := bson.Marshal(bson.M{"data": value})
	return len(v)
}
