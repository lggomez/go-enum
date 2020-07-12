package templates

var EnumCodecsTestTemplate = `{{- if not .OmitGeneratedNotice}}
// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// {{ .Timestamp }}
// by go-enum
{{- end }}

package {{ .Package }}

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

type fooEnumValue struct {
	EnumValue stringEnumValue ` + "`" + `json:"enum_value"` + "`" + `
}

type fooEnumValueOmitEmpty struct {
	EnumValue stringEnumValue ` + "`" + `json:"enum_value,omitempty"` + "`" + `
}

type fooEnumValuePtr struct {
	EnumValue *stringEnumValue ` + "`" + `json:"enum_value"` + "`" + `
}

type fooEnumValuePtrOmitEmpty struct {
	EnumValue *stringEnumValue ` + "`" + `json:"enum_value,omitempty"` + "`" + `
}

func TestEnumValue_MarshalJSON(t *testing.T) {
	t.Run("Marshal_AnnonStructField", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }
		v := struct {
			A stringEnumValue ` + "`" + `json:"enum_value"` + "`" + `
		}{c}
		data, err := json.Marshal(&v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(` + "`" + `{"enum_value":"%s"}` + "`" + `, v.A.String()), string(data))
	})
	t.Run("Marshal_StructField", func(t *testing.T) {
		v := fooEnumValue{EnumValue: stringEnumValue{key: {{ .IndexKeyName }} }}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(` + "`" + `{"enum_value":"%s"}` + "`" + `, v.EnumValue.String()), string(data))
	})
	t.Run("Marshal_OmitEmptyStruct", func(t *testing.T) {
		// encoding/json ignores omitempty on struct zero values
		// https://github.com/golang/go/issues/11939
		v := fooEnumValueOmitEmpty{}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, ` + "`" + `{"enum_value":""}` + "`" + `, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		vPtr := fooEnumValuePtr{EnumValue: &stringEnumValue{key: {{ .IndexKeyName }} }}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, ` + "`" + `{"enum_value":""}` + "`" + `, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }
		vPtr := fooEnumValuePtr{EnumValue: &c}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(` + "`" + `{"enum_value":"%s"}` + "`" + `, vPtr.EnumValue.String()), string(data))
	})
	t.Run("Marshal_StructFieldNilPtr", func(t *testing.T) {
		vPtr := fooEnumValuePtr{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, ` + "`" + `{"enum_value":null}` + "`" + `, string(data))
	})
	t.Run("Marshal_OmitEmptyStructPtr", func(t *testing.T) {
		vPtr := fooEnumValuePtrOmitEmpty{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, ` + "`" + `{}` + "`" + `, string(data))
	})
}

func TestEnumValue_UnmarshalJSON(t *testing.T) {
	t.Run("Unmarshal_InvalidJSON", func(t *testing.T) {
		data := ` + "`" + `{"enum_value":"PA` + "`" + `
rawData := []byte(data)

c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }
v := struct {
A stringEnumValue ` + "`" + `json:"enum_value"` + "`" + `
}{c}
// Invalid JSON structures fail on json.Unmarshal
err := json.Unmarshal(rawData, &v)
require.NotNil(t, err)
})
t.Run("Unmarshal_InvalidValue", func(t *testing.T) {
	data := ` + "`" + `{"enum_value":"{{ .TestCaseInvalidValue }}"}` + "`" + `
	rawData := []byte(data)

	c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }
	v := struct {
		A stringEnumValue ` + "`" + `json:"enum_value"` + "`" + `
	}{c}
	err := json.Unmarshal(rawData, &v)
	require.NotNil(t, err)
})
t.Run("Unmarshal_InvalidValueJSON", func(t *testing.T) {
	data := ` + "`" + `{"enum_value":123}` + "`" + `
	rawData := []byte(data)

	c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }
	v := struct {
		A stringEnumValue ` + "`" + `json:"enum_value"` + "`" + `
	}{c}
	// Invalid field values but whithin a valid JSON are caught on UnmarshalJSON
	err := json.Unmarshal(rawData, &v)
	require.NotNil(t, err)
})
t.Run("Unmarshal_AnnonStructField", func(t *testing.T) {
	data := ` + "`" + `{"enum_value":"{{ .TestCaseValue }}"}` + "`" + `
	rawData := []byte(data)

	c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }
	v := struct {
		A stringEnumValue ` + "`" + `json:"enum_value"` + "`" + `
	}{c}
	err := json.Unmarshal(rawData, &v)
	require.Nil(t, err)
	assert.EqualValues(t, "{{ .TestCaseValue }}", v.A.String())
})
t.Run("Unmarshal_StructField", func(t *testing.T) {
	data := ` + "`" + `{"enum_value":"{{ .TestCaseValue }}"}` + "`" + `
	rawData := []byte(data)

	v := fooEnumValue{EnumValue: stringEnumValue{key: {{ .IndexKeyName }} }}
	err := json.Unmarshal(rawData, &v)
	require.Nil(t, err)
	assert.EqualValues(t, "{{ .TestCaseValue }}", v.EnumValue.String())
})
t.Run("Unmarshal_OmitEmptyStruct", func(t *testing.T) {
	data := ` + "`" + `{"enum_value":null}` + "`" + `
	rawData := []byte(data)

	v := fooEnumValueOmitEmpty{}
	err := json.Unmarshal(rawData, &v)
	require.NoError(t, err)
})
t.Run("Unmarshal_StructFieldPtr", func(t *testing.T) {
	data := ` + "`" + `{"enum_value":"{{ .TestCaseValue }}"}` + "`" + `
	rawData := []byte(data)

	vPtr := fooEnumValuePtr{EnumValue: &stringEnumValue{key: {{ .IndexKeyName }} }}
	err := json.Unmarshal(rawData, &vPtr)
	require.Nil(t, err)
	assert.EqualValues(t, "{{ .TestCaseValue }}", vPtr.EnumValue.String())
})
t.Run("Unmarshal_StructFieldNilPtr", func(t *testing.T) {
	data := ` + "`" + `{"enum_value":null}` + "`" + `
	rawData := []byte(data)

	vPtr := fooEnumValuePtr{}
	err := json.Unmarshal(rawData, &vPtr)
	require.Nil(t, err)
	assert.Nil(t, vPtr.EnumValue)
})
t.Run("Unmarshal_TableUnmatch", func(t *testing.T) {
	data := ` + "`" + `{"enum_value":"PPP"}` + "`" + `
	rawData := []byte(data)

	v := fooEnumValue{}
	err := json.Unmarshal(rawData, &v)
	require.NotNil(t, err)
})
}

func TestEnumValue_TextCodec(t *testing.T) {
	t.Run("MarshalText_Valid", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }

		data, err := c.MarshalText()
		require.Nil(t, err)
		require.True(t, len(data) > 0)
	})
	t.Run("UnmarshalText_Valid", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		err := c.UnmarshalText([]byte("{{ .TestCaseValue }}"))
		require.Nil(t, err)
		require.EqualValues(t, {{ .TestCaseName }}.String(), c.value)
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		err := c.UnmarshalText([]byte(""))
		require.NoError(t, err)
		err = c.UnmarshalText([]byte("{{ .TestCaseInvalidValue }}"))
		require.NotNil(t, err)
	})
}

func TestEnumValue_Stringer(t *testing.T) {
	c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }
	require.EqualValues(t, c.value, c.String())
}

func TestEnumValue_DriverValues(t *testing.T) {
	t.Run("Scan_String", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		require.Nil(t, c.Scan({{ .TestCaseName }}.String()))
		require.EqualValues(t, {{ .TestCaseName }}.String(), c.value)
	})
	t.Run("Scan_Bytes", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		require.Nil(t, c.Scan([]byte({{ .TestCaseName }}.String())))
		require.EqualValues(t, {{ .TestCaseName }}.String(), c.value)
	})
	t.Run("Scan_Invalid", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		require.NotNil(t, c.Scan(1))
	})
	t.Run("Scan_Invalid", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }

		v, err := c.Value()
		assert.Nil(t, err)
		require.EqualValues(t, {{ .TestCaseName }}.String(), fmt.Sprintf("%v", v))
	})
}

func TestEnumValue_BinaryCodec(t *testing.T) {
	t.Run("MarshalBinary_Valid", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }

		data, err := c.MarshalBinary()
		require.Nil(t, err)
		assert.Len(t, data, {{ .TestCaseBinaryLen }})
	})
	t.Run("UnmarshalBinary_Valid", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		err := c.UnmarshalBinary([]byte("{{ .TestCaseValue }}"))
		require.Nil(t, err)
		require.EqualValues(t, {{ .TestCaseName }}.String(), c.value)
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		err := c.UnmarshalBinary(nil)
		require.NoError(t, err)
		err = c.UnmarshalBinary([]byte("{{ .TestCaseInvalidValue }}"))
		require.NotNil(t, err)
	})
}

func TestEnumValue_GobCodec(t *testing.T) {
	t.Run("MarshalBinary_Valid", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }

		data, err := c.GobEncode()
		require.Nil(t, err)
		assert.Len(t, data, {{ .TestCaseBinaryLen }})
	})
	t.Run("UnmarshalBinary_Valid", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		err := c.GobDecode([]byte("{{ .TestCaseValue }}"))
		require.Nil(t, err)
		require.EqualValues(t, {{ .TestCaseName }}.String(), c.value)
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &stringEnumValue{key: {{ .IndexKeyName }} }
		err := c.GobDecode(nil)
		require.NoError(t, err)
		err = c.GobDecode([]byte("{{ .TestCaseInvalidValue }}"))
		require.NotNil(t, err)
	})
}

func TestEnumValue_MarshalBSON(t *testing.T) {
	t.Run("MarshalBSON_Valid", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }

		data, err := c.MarshalBSON()
		require.Nil(t, err)
		assert.Len(t, data, {{ .TestCaseBSONLen }})
	})
}

func TestEnumValue_UnmarshalBSON(t *testing.T) {
	t.Run("UnmarshalBSON_Valid", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }

		v1 := &fooEnumValue{EnumValue: c}
		rawData, err := bson.Marshal(v1)
		require.Nil(t, err)

		v2 := &fooEnumValue{EnumValue: c}
		err = bson.Unmarshal(rawData, &v2)
		require.Nil(t, err)
		assert.EqualValues(t, "{{ .TestCaseValue }}", v2.EnumValue.String())
	})
	t.Run("UnmarshalBSON_InvalidTable", func(t *testing.T) {
		c := stringEnumValue{ {{ .TestCaseName }}.String(), {{ .IndexKeyName }} }
		ptr := &c
		ptr.value = "PPP"
		v1 := &fooEnumValue{EnumValue: c}
		rawData, err := bson.Marshal(v1)
		require.Nil(t, err)

		v2 := &fooEnumValue{}
		err = bson.Unmarshal(rawData, &v2)
		require.NotNil(t, err)
	})
}

`
