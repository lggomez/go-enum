package gen

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/stoewer/go-strcase"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/lggomez/go-enum/gen/internal"
	"github.com/lggomez/go-enum/gen/internal/templates"
)

type StringEnumDefinition struct {
	Name   string
	Values []string
}

type canonicalStringEnum struct {
	StructName          string
	StructNameLowerCase string

	IndexKeyName string
	Values       map[string]string

	TestCaseName         string
	TestCaseKey          string
	TestCaseValue        string
	TestCaseInvalidValue string
	TestCaseBinaryLen    int
	TestCaseBSONLen      int

	ImportPath          string
	FileName            string
	Timestamp           string
	Package             string
	OmitGeneratedNotice bool
}

func GenerateEnumTypes(packageFilePath string, packageImportPath string, omitGeneratedNotice bool, enums ...StringEnumDefinition) {
	// Get package name from import path
	// i.e: github.com/lggomez/go-enum/example -> example
	tokens := strings.Split(packageImportPath, "/")
	packageName := tokens[len(tokens)-1]

	// Convert enum definitions into canonical definitions with full metadata for code generation
	canonicalEnums := processEnumerations(packageImportPath, packageName, omitGeneratedNotice, enums)

	if _, err := os.Stat(packageFilePath); os.IsNotExist(err) {
		if dirErr := os.Mkdir(packageFilePath, os.ModePerm); dirErr != nil {
			log.Panic("could not create package - ", err.Error())
		}
	}

	// Traverse enums for code generation
	for i, canonicalEnum := range canonicalEnums {
		// Generate base enum struct and its codecs
		// This is a single time pass that must be done on the first iteration
		if i == 0 {
			if err := generateFileFromTemplate(canonicalEnum,
				templates.EnumTemplate,
				fmt.Sprintf("%s%senum.go", packageFilePath, string(os.PathSeparator))); err != nil {
				log.Panic(err.Error())
			}

			if err := generateFileFromTemplate(canonicalEnum,
				templates.EnumCodecsTemplate,
				fmt.Sprintf("%s%senum_codecs.go", packageFilePath, string(os.PathSeparator))); err != nil {
				log.Panic(err.Error())
			}
		}

		// Generate specific enum implementation file
		if err := generateFileFromTemplate(canonicalEnum,
			templates.EnumImplTemplate,
			fmt.Sprintf("%s%s%s.go", packageFilePath, string(os.PathSeparator), canonicalEnum.FileName)); err != nil {
			log.Panic(err.Error())
		}

		// Generate specific enum implementation test file
		if err := generateFileFromTemplate(canonicalEnum,
			templates.EnumImplTestTemplate,
			fmt.Sprintf("%s%s%s_test.go", packageFilePath, string(os.PathSeparator), canonicalEnum.FileName)); err != nil {
			log.Panic(err.Error())
		}

		// Generate codecs test file
		// This is a single time pass that must be done on the first iteration,
		// and after the first specific enum is generated since it uses it for tests
		if i == 0 {
			if err := generateFileFromTemplate(canonicalEnum,
				templates.EnumCodecsTestTemplate,
				fmt.Sprintf("%s%senum_codecs_test.go", packageFilePath, string(os.PathSeparator))); err != nil {
				log.Panic(err.Error())
			}
		}
	}
}

func generateFileFromTemplate(canonicalEnum canonicalStringEnum, templateString, destinationPath string) error {
	src := &bytes.Buffer{}
	enumTemplate := template.Must(template.New(destinationPath).Parse(templateString))
	err := enumTemplate.Execute(src, canonicalEnum)
	if err != nil {
		log.Panic(err.Error())
	}

	formattedSrc, err := internal.FormatSource(src)
	if err != nil {
		log.Panic(err.Error())
	}

	err = internal.SaveFile(destinationPath, formattedSrc)
	if err != nil {
		log.Panic(err.Error())
	}

	return nil
}

func processEnumerations(importPath string, packageName string, omitGeneratedNotice bool, enums []StringEnumDefinition) []canonicalStringEnum {
	canonicalEnums := make([]canonicalStringEnum, len(enums))

	for i, e := range enums {
		upperCamelName := strcase.UpperCamelCase(e.Name)
		lowerCamelName := strcase.LowerCamelCase(e.Name)

		ce := &canonicalStringEnum{
			Package:              packageName,
			ImportPath:           importPath,
			StructName:           upperCamelName,
			StructNameLowerCase:  lowerCamelName,
			IndexKeyName:         lowerCamelName + "Key",
			Values:               map[string]string{},
			TestCaseKey:          lowerCamelName,
			TestCaseInvalidValue: uuid.New().String(), // Set an unique random value to prevent collisions
			FileName:             strcase.SnakeCase(e.Name),
			Timestamp:            time.Now().String(),
			OmitGeneratedNotice:  omitGeneratedNotice,
		}

		for i, value := range e.Values {
			valueKey := strcase.UpperCamelCase(value)
			if i == 0 {
				ce.TestCaseName = ce.StructName + valueKey
				ce.TestCaseKey = ce.StructName
				ce.TestCaseValue = value
				ce.TestCaseBinaryLen = len([]byte(ce.TestCaseValue))
				ce.TestCaseBSONLen = calculateBSONLen(ce.TestCaseValue)
			}
			ce.Values[valueKey] = value
		}

		canonicalEnums[i] = *ce
	}

	return canonicalEnums
}

func calculateBSONLen(value string) int {
	v, _ := bson.Marshal(bson.M{"data": value})
	return len(v)
}
