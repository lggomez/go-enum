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

	IndexKeyName  string
	IndexKeyValue string
	Values        map[string]string

	TestCaseName         string
	TestCaseKey          string
	TestCaseValue        string
	TestCaseInvalidValue string
	TestCaseBinaryLen    int
	TestCaseBSONLen      int

	ImportPath string
	FileName   string
	Timestamp  string
	Package    string
}

func GenerateEnumTypes(currentPackagePath string, enums ...StringEnumDefinition) {
	path := "."
	tokens := strings.Split(currentPackagePath, "/")
	packageName := tokens[len(tokens)-1]
	canonicalEnums := processEnumerations(currentPackagePath, packageName, enums)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Panic(err.Error())
	}

	for i, canonicalEnum := range canonicalEnums {
		if i == 0 {
			if err := generateFileFromTemplate(canonicalEnum,
				templates.EnumTemplate,
				fmt.Sprintf("%s%senum.go", path, string(os.PathSeparator))); err != nil {
				log.Panic(err.Error())
			}

			if err := generateFileFromTemplate(canonicalEnum,
				templates.EnumCodecsTemplate,
				fmt.Sprintf("%s%senum_codecs.go", path, string(os.PathSeparator))); err != nil {
				log.Panic(err.Error())
			}
		}

		if err := generateFileFromTemplate(canonicalEnum,
			templates.EnumImplTemplate,
			fmt.Sprintf("%s%s%s.go", path, string(os.PathSeparator), canonicalEnum.FileName)); err != nil {
			log.Panic(err.Error())
		}

		if err := generateFileFromTemplate(canonicalEnum,
			templates.EnumImplTestTemplate,
			fmt.Sprintf("%s%s%s_test.go", path, string(os.PathSeparator), canonicalEnum.FileName)); err != nil {
			log.Panic(err.Error())
		}

		if i == 0 {
			if err := generateFileFromTemplate(canonicalEnum,
				templates.EnumCodecsTestTemplate,
				fmt.Sprintf("%s%senum_codecs_test.go", path, string(os.PathSeparator))); err != nil {
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

func processEnumerations(importPath string, packageName string, enums []StringEnumDefinition) []canonicalStringEnum {
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
			IndexKeyValue:        upperCamelName,
			Values:               map[string]string{},
			TestCaseKey:          lowerCamelName,
			TestCaseInvalidValue: uuid.New().String(),
			FileName:             strcase.SnakeCase(e.Name),
			Timestamp:            time.Now().String(),
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
