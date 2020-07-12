package templates

var EnumImplTestTemplate = `{{- if not .OmitGeneratedNotice}}
// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// {{ .Timestamp }}
// by go-enum
{{- end }}

package {{ .Package }}_test

import (
	"testing"
	"encoding/json"
	"fmt"

	"{{ .ImportPath }}"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type foo{{ .StructName }} struct {
	TestValue {{ .Package }}.{{ .StructName }} ` + "`" + `json:"{{ .TestCaseKey }}"` + "`" + `
}

type foo{{ .StructName }}OmitEmpty struct {
	TestValue {{ .Package }}.{{ .StructName }} ` + "`" + `json:"{{ .TestCaseKey }},omitempty"` + "`" + `
}

type foo{{ .StructName }}Ptr struct {
	TestValue *{{ .Package }}.{{ .StructName }} ` + "`" + `json:"{{ .TestCaseKey }}"` + "`" + `
}

type foo{{ .StructName }}PtrOmitEmpty struct {
	TestValue *{{ .Package }}.{{ .StructName }} ` + "`" + `json:"{{ .TestCaseKey }},omitempty"` + "`" + `
}

func Test{{ .StructName }}_MarshalJSON(t *testing.T) {
	t.Run("Marshal_AnnonStructField", func(t *testing.T) {
		v := struct {
			A {{ .Package }}.{{ .StructName }} ` + "`" + `json:"{{ .TestCaseKey }}"` + "`" + `
		}{ {{ .Package }}.{{ .TestCaseName }}}
		data, err := json.Marshal(&v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(` + "`" + `{"{{ .TestCaseKey }}":"%s"}` + "`" + `, v.A.String()), string(data))
	})
	t.Run("Marshal_StructField", func(t *testing.T) {
		v := foo{{ .StructName }}{TestValue: {{ .Package }}.{{ .StructName }}{}}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(` + "`" + `{"{{ .TestCaseKey }}":"%s"}` + "`" + `, v.TestValue.String()), string(data))
	})
	t.Run("Marshal_OmitEmptyStruct", func(t *testing.T) {
		// encoding/json ignores omitempty on struct zero values
		// https://github.com/golang/go/issues/11939
		v := foo{{ .StructName }}OmitEmpty{}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, ` + "`" + `{"{{ .TestCaseKey }}":""}` + "`" + `, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		vPtr := foo{{ .StructName }}Ptr{TestValue: &{{ .Package }}.{{ .StructName }}{}}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, ` + "`" + `{"{{ .TestCaseKey }}":""}` + "`" + `, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		dt := {{ .Package }}.{{ .TestCaseName }}
		vPtr := foo{{ .StructName }}Ptr{TestValue: &dt}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(` + "`" + `{"{{ .TestCaseKey }}":"%s"}` + "`" + `, vPtr.TestValue.String()), string(data))
	})
	t.Run("Marshal_StructFieldNilPtr", func(t *testing.T) {
		vPtr := foo{{ .StructName }}Ptr{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, ` + "`" + `{"{{ .TestCaseKey }}":null}` + "`" + `, string(data))
	})
	t.Run("Marshal_OmitEmptyStructPtr", func(t *testing.T) {
		vPtr := foo{{ .StructName }}PtrOmitEmpty{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, ` + "`" + `{}` + "`" + `, string(data))
	})
}

func Test{{ .StructName }}_UnmarshalJSON(t *testing.T) {
	t.Run("Unmarshal_InvalidValue", func(t *testing.T) {
		data := ` + "`" + `{"{{ .TestCaseKey }}":"{{ .TestCaseInvalidValue }}"}` + "`" + `
		rawData := []byte(data)

		v := struct {
			A {{ .Package }}.{{ .StructName }} ` + "`" + `json:"{{ .TestCaseKey }}"` + "`" + `
		}{ {{ .Package }}.{{ .TestCaseName }}}
		err := json.Unmarshal(rawData, &v)
		require.NotNil(t, err)
	})
	t.Run("Unmarshal_AnnonStructField", func(t *testing.T) {
		data := ` + "`" + `{"{{ .TestCaseKey }}":"{{ .TestCaseValue }}"}` + "`" + `
		rawData := []byte(data)

		v := struct {
			A {{ .Package }}.{{ .StructName }} ` + "`" + `json:"{{ .TestCaseKey }}"` + "`" + `
		}{ {{ .Package }}.{{ .TestCaseName }} }
		err := json.Unmarshal(rawData, &v)
		require.Nil(t, err)
		assert.EqualValues(t, "{{ .TestCaseValue }}", v.A.String())
	})
	t.Run("Unmarshal_StructField", func(t *testing.T) {
		data := ` + "`" + `{"{{ .TestCaseKey }}":"{{ .TestCaseValue }}"}` + "`" + `
		rawData := []byte(data)

		v := foo{{ .StructName }}{}
		err := json.Unmarshal(rawData, &v)
		require.Nil(t, err)
		assert.EqualValues(t, "{{ .TestCaseValue }}", v.TestValue.String())
	})
	t.Run("Unmarshal_OmitEmptyStruct", func(t *testing.T) {
		data := ` + "`" + `{"{{ .TestCaseKey }}":null}` + "`" + `
		rawData := []byte(data)

		v := foo{{ .StructName }}OmitEmpty{}
		err := json.Unmarshal(rawData, &v)
		require.NoError(t, err)
	})
	t.Run("Unmarshal_StructFieldPtr", func(t *testing.T) {
		data := ` + "`" + `{"{{ .TestCaseKey }}":"{{ .TestCaseValue }}"}` + "`" + `
		rawData := []byte(data)

		vPtr := foo{{ .StructName }}Ptr{}
		err := json.Unmarshal(rawData, &vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, "{{ .TestCaseValue }}", vPtr.TestValue.String())
	})
	t.Run("Unmarshal_StructFieldNilPtr", func(t *testing.T) {
		data := ` + "`" + `{"{{ .TestCaseKey }}":null}` + "`" + `
		rawData := []byte(data)

		vPtr := foo{{ .StructName }}Ptr{}
		err := json.Unmarshal(rawData, &vPtr)
		require.Nil(t, err)
		assert.Nil(t, vPtr.TestValue)
	})
}

func Test{{ .StructName }}_EmptyValues(t *testing.T) {
	t.Run("StructField", func(t *testing.T) {
		v := foo{{ .StructName }}{TestValue: {{ .Package }}.{{ .StructName }}{}}
		data, err := json.Marshal(v)
		require.NoError(t, err)
		assert.EqualValues(t, fmt.Sprintf(` + "`" + `{"{{ .TestCaseKey }}":"%s"}` + "`" + `, v.TestValue.String()), string(data))

		v2 := foo{{ .StructName }}{}
		err = json.Unmarshal(data, &v2)
		require.NoError(t, err) // empty field

		v3 := foo{{ .StructName }}{}
		err = json.Unmarshal(data, &v3)
		require.NoError(t, err)
	})
	t.Run("PtrField", func(t *testing.T) {
		v := foo{{ .StructName }}Ptr{TestValue: nil}
		data, err := json.Marshal(v)
		require.NoError(t, err)
		assert.EqualValues(t, fmt.Sprintf(` + "`" + `{"{{ .TestCaseKey }}":%s}` + "`" + `, "null"), string(data))

		v2 := foo{{ .StructName }}{}
		err = json.Unmarshal(data, &v2)
		require.NoError(t, err) // empty field

		v3 := foo{{ .StructName }}{}
		err = json.Unmarshal(data, &v3)
		require.NoError(t, err)
	})
}

func Test{{ .StructName }}_ForEach(t *testing.T) {
	j := 0
	{{ .Package }}.Enum{{ .StructName }}.ForEach(func(i int, enumValue {{ .Package }}.{{ .StructName }}) {
		j++
		value, found := {{ .Package }}.{{ .StructName }}FromValue(enumValue.String(), false)
		assert.True(t, found)
		assert.True(t, value == enumValue)
	})
	assert.EqualValues(t, {{ .Package }}.Enum{{ .StructName }}.Len(), j)
}

`
