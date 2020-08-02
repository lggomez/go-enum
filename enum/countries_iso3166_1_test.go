// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2020-08-02T15:36:10-03:00
// by go-enum

package enum_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lggomez/go-enum/example/enum"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

type fooCountriesIso31661 struct {
	TestValue enum.CountriesIso31661 `json:"countries_iso3166_1"`
}

type fooCountriesIso31661OmitEmpty struct {
	TestValue enum.CountriesIso31661 `json:"countries_iso3166_1,omitempty"`
}

type fooCountriesIso31661Ptr struct {
	TestValue *enum.CountriesIso31661 `json:"countries_iso3166_1"`
}

type fooCountriesIso31661PtrOmitEmpty struct {
	TestValue *enum.CountriesIso31661 `json:"countries_iso3166_1,omitempty"`
}

func TestCountriesIso31661_MarshalJSON(t *testing.T) {
	t.Run("Marshal_AnnonStructField", func(t *testing.T) {
		v := struct {
			A enum.CountriesIso31661 `json:"countries_iso3166_1"`
		}{enum.CountriesIso31661CA}
		data, err := json.Marshal(&v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"countries_iso3166_1":"%s"}`, v.A.String()), string(data))
	})
	t.Run("Marshal_StructField", func(t *testing.T) {
		v := fooCountriesIso31661{TestValue: enum.CountriesIso31661{}}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"countries_iso3166_1":"%s"}`, v.TestValue.String()), string(data))
	})
	t.Run("Marshal_OmitEmptyStruct", func(t *testing.T) {
		// encoding/json ignores omitempty on struct zero values
		// https://github.com/golang/go/issues/11939
		v := fooCountriesIso31661OmitEmpty{}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, `{"countries_iso3166_1":""}`, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		vPtr := fooCountriesIso31661Ptr{TestValue: &enum.CountriesIso31661{}}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{"countries_iso3166_1":""}`, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		dt := enum.CountriesIso31661CA
		vPtr := fooCountriesIso31661Ptr{TestValue: &dt}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"countries_iso3166_1":"%s"}`, vPtr.TestValue.String()), string(data))
	})
	t.Run("Marshal_StructFieldNilPtr", func(t *testing.T) {
		vPtr := fooCountriesIso31661Ptr{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{"countries_iso3166_1":null}`, string(data))
	})
	t.Run("Marshal_OmitEmptyStructPtr", func(t *testing.T) {
		vPtr := fooCountriesIso31661PtrOmitEmpty{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{}`, string(data))
	})
}

func TestCountriesIso31661_UnmarshalJSON(t *testing.T) {
	t.Run("Unmarshal_InvalidValue", func(t *testing.T) {
		data := `{"countries_iso3166_1":"invalid_e55bd60f-6352-4861-a0be-048acf7c3d03"}`
		rawData := []byte(data)

		v := struct {
			A enum.CountriesIso31661 `json:"countries_iso3166_1"`
		}{enum.CountriesIso31661CA}
		err := json.Unmarshal(rawData, &v)
		require.NotNil(t, err)
	})
	t.Run("Unmarshal_AnnonStructField", func(t *testing.T) {
		data := `{"countries_iso3166_1":"Ca"}`
		rawData := []byte(data)

		v := struct {
			A enum.CountriesIso31661 `json:"countries_iso3166_1"`
		}{enum.CountriesIso31661CA}
		err := json.Unmarshal(rawData, &v)
		require.Nil(t, err)
		assert.EqualValues(t, "Ca", v.A.String())
	})
	t.Run("Unmarshal_StructField", func(t *testing.T) {
		data := `{"countries_iso3166_1":"Ca"}`
		rawData := []byte(data)

		v := fooCountriesIso31661{}
		err := json.Unmarshal(rawData, &v)
		require.Nil(t, err)
		assert.EqualValues(t, "Ca", v.TestValue.String())
	})
	t.Run("Unmarshal_OmitEmptyStruct", func(t *testing.T) {
		data := `{"countries_iso3166_1":null}`
		rawData := []byte(data)

		v := fooCountriesIso31661OmitEmpty{}
		err := json.Unmarshal(rawData, &v)
		require.NoError(t, err)
	})
	t.Run("Unmarshal_StructFieldPtr", func(t *testing.T) {
		data := `{"countries_iso3166_1":"Ca"}`
		rawData := []byte(data)

		vPtr := fooCountriesIso31661Ptr{}
		err := json.Unmarshal(rawData, &vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, "Ca", vPtr.TestValue.String())
	})
	t.Run("Unmarshal_StructFieldNilPtr", func(t *testing.T) {
		data := `{"countries_iso3166_1":null}`
		rawData := []byte(data)

		vPtr := fooCountriesIso31661Ptr{}
		err := json.Unmarshal(rawData, &vPtr)
		require.Nil(t, err)
		assert.Nil(t, vPtr.TestValue)
	})
}

func TestCountriesIso31661_EmptyValues(t *testing.T) {
	t.Run("StructField", func(t *testing.T) {
		v := fooCountriesIso31661{TestValue: enum.CountriesIso31661{}}
		data, err := json.Marshal(v)
		require.NoError(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"countries_iso3166_1":"%s"}`, v.TestValue.String()), string(data))

		v2 := fooCountriesIso31661{}
		err = json.Unmarshal(data, &v2)
		require.NoError(t, err)
		require.True(t, v2.TestValue.IsEmpty())

		v3 := fooCountriesIso31661{}
		err = json.Unmarshal(data, &v3)
		require.NoError(t, err)
		require.True(t, v3.TestValue.IsEmpty())
	})
	t.Run("PtrField", func(t *testing.T) {
		v := fooCountriesIso31661Ptr{TestValue: nil}
		data, err := json.Marshal(v)
		require.NoError(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"countries_iso3166_1":%s}`, "null"), string(data))

		v2 := fooCountriesIso31661{}
		err = json.Unmarshal(data, &v2)
		require.NoError(t, err)
		require.True(t, v2.TestValue.IsEmpty())

		v3 := fooCountriesIso31661{}
		err = json.Unmarshal(data, &v3)
		require.NoError(t, err)
		require.True(t, v3.TestValue.IsEmpty())
	})
}

func TestCountriesIso31661_UndefinedValue(t *testing.T) {
	require.True(t, enum.CountriesIso31661{}.IsUndefined())
}

func TestCountriesIso31661_ForEach(t *testing.T) {
	j := 0
	enum.EnumCountriesIso31661.ForEach(func(i int, enumValue enum.CountriesIso31661) {
		j++
		value, found := enum.CountriesIso31661FromValue(enumValue.String(), false)
		assert.True(t, found)
		assert.True(t, value == enumValue)

		assert.True(t, enumValue.Equals(value.String()))
		assert.True(t, enumValue.EqualsIgnoreCase(value.String()))
		assert.True(t, value.Equals(enumValue.String()))
		assert.True(t, value.EqualsIgnoreCase(enumValue.String()))
	})
	assert.EqualValues(t, enum.EnumCountriesIso31661.Len(), j)
}

func TestCountriesIso31661_TextCodec(t *testing.T) {
	t.Run("MarshalText_Valid", func(t *testing.T) {
		c := fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}

		data, err := c.TestValue.MarshalText()
		require.Nil(t, err)
		require.True(t, len(data) > 0)
	})
	t.Run("UnmarshalText_Valid", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		err := c.TestValue.UnmarshalText([]byte("Ca"))
		require.Nil(t, err)
		require.EqualValues(t, enum.CountriesIso31661CA.String(), c.TestValue.String())
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		err := c.TestValue.UnmarshalText([]byte(""))
		require.NoError(t, err)
		err = c.TestValue.UnmarshalText([]byte("invalid_e55bd60f-6352-4861-a0be-048acf7c3d03"))
		require.NotNil(t, err)
	})
}

func TestCountriesIso31661_Stringer(t *testing.T) {
	c := fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
	require.EqualValues(t, "Ca", c.TestValue.String())
}

func TestCountriesIso31661_DriverValues(t *testing.T) {
	t.Run("Scan_String", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		require.Nil(t, c.TestValue.Scan(enum.CountriesIso31661CA.String()))
		require.EqualValues(t, enum.CountriesIso31661CA.String(), c.TestValue.String())
	})
	t.Run("Scan_Bytes", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		require.Nil(t, c.TestValue.Scan([]byte(enum.CountriesIso31661CA.String())))
		require.EqualValues(t, enum.CountriesIso31661CA.String(), c.TestValue.String())
	})
	t.Run("Scan_Invalid", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		require.NotNil(t, c.TestValue.Scan(1))
	})
	t.Run("Scan_Invalid", func(t *testing.T) {
		c := fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}

		v, err := c.TestValue.Value()
		assert.Nil(t, err)
		require.EqualValues(t, enum.CountriesIso31661CA.String(), fmt.Sprintf("%v", v))
	})
}

func TestCountriesIso31661_BinaryCodec(t *testing.T) {
	t.Run("MarshalBinary_Valid", func(t *testing.T) {
		c := fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}

		data, err := c.TestValue.MarshalBinary()
		require.Nil(t, err)
		assert.Len(t, data, 2)
	})
	t.Run("UnmarshalBinary_Valid", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		err := c.TestValue.UnmarshalBinary([]byte("Ca"))
		require.Nil(t, err)
		require.EqualValues(t, enum.CountriesIso31661CA.String(), c.TestValue.String())
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		err := c.TestValue.UnmarshalBinary(nil)
		require.NoError(t, err)
		err = c.TestValue.UnmarshalBinary([]byte("invalid_e55bd60f-6352-4861-a0be-048acf7c3d03"))
		require.NotNil(t, err)
	})
}

func TestCountriesIso31661_GobCodec(t *testing.T) {
	t.Run("MarshalBinary_Valid", func(t *testing.T) {
		c := fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}

		data, err := c.TestValue.GobEncode()
		require.Nil(t, err)
		assert.Len(t, data, 2)
	})
	t.Run("UnmarshalBinary_Valid", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		err := c.TestValue.GobDecode([]byte("Ca"))
		require.Nil(t, err)
		require.EqualValues(t, enum.CountriesIso31661CA.String(), c.TestValue.String())
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}
		err := c.TestValue.GobDecode(nil)
		require.NoError(t, err)
		err = c.TestValue.GobDecode([]byte("invalid_e55bd60f-6352-4861-a0be-048acf7c3d03"))
		require.NotNil(t, err)
	})
}

func TestCountriesIso31661_MarshalBSON(t *testing.T) {
	t.Run("MarshalBSON_Valid", func(t *testing.T) {
		c := fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}

		data, err := c.TestValue.MarshalBSON()
		require.Nil(t, err)
		assert.Len(t, data, 18)
	})
}

func TestCountriesIso31661_UnmarshalBSON(t *testing.T) {
	t.Run("UnmarshalBSON_Valid", func(t *testing.T) {
		c := fooCountriesIso31661{TestValue: enum.CountriesIso31661CA}

		v1 := &c
		rawData, err := bson.Marshal(v1)
		require.Nil(t, err)

		v2 := &fooCountriesIso31661{}
		err = bson.Unmarshal(rawData, &v2)
		require.Nil(t, err)
		assert.EqualValues(t, "Ca", v2.TestValue.String())
	})
}
