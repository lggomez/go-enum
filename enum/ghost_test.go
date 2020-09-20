// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2020-08-02T15:35:34-03:00
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

type fooGhost struct {
	TestValue enum.Ghost `json:"ghost"`
}

type fooGhostOmitEmpty struct {
	TestValue enum.Ghost `json:"ghost,omitempty"`
}

type fooGhostPtr struct {
	TestValue *enum.Ghost `json:"ghost"`
}

type fooGhostPtrOmitEmpty struct {
	TestValue *enum.Ghost `json:"ghost,omitempty"`
}

func TestGhost_MarshalJSON(t *testing.T) {
	t.Run("Marshal_AnnonStructField", func(t *testing.T) {
		v := struct {
			A enum.Ghost `json:"ghost"`
		}{enum.GhostBlinky}
		data, err := json.Marshal(&v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"ghost":"%s"}`, v.A.String()), string(data))
	})
	t.Run("Marshal_StructField", func(t *testing.T) {
		v := fooGhost{TestValue: enum.Ghost{}}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"ghost":"%s"}`, v.TestValue.String()), string(data))
	})
	t.Run("Marshal_OmitEmptyStruct", func(t *testing.T) {
		// encoding/json ignores omitempty on struct zero values
		// https://github.com/golang/go/issues/11939
		v := fooGhostOmitEmpty{}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, `{"ghost":""}`, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		vPtr := fooGhostPtr{TestValue: &enum.Ghost{}}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{"ghost":""}`, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		dt := enum.GhostBlinky
		vPtr := fooGhostPtr{TestValue: &dt}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"ghost":"%s"}`, vPtr.TestValue.String()), string(data))
	})
	t.Run("Marshal_StructFieldNilPtr", func(t *testing.T) {
		vPtr := fooGhostPtr{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{"ghost":null}`, string(data))
	})
	t.Run("Marshal_OmitEmptyStructPtr", func(t *testing.T) {
		vPtr := fooGhostPtrOmitEmpty{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{}`, string(data))
	})
}

func TestGhost_UnmarshalJSON(t *testing.T) {
	t.Run("Unmarshal_InvalidValue", func(t *testing.T) {
		data := `{"ghost":"invalid_e6831993-894b-4957-ac84-c503d35e1a17"}`
		rawData := []byte(data)

		v := struct {
			A enum.Ghost `json:"ghost"`
		}{enum.GhostBlinky}
		err := json.Unmarshal(rawData, &v)
		require.NotNil(t, err)
	})
	t.Run("Unmarshal_AnnonStructField", func(t *testing.T) {
		data := `{"ghost":"Blinky"}`
		rawData := []byte(data)

		v := struct {
			A enum.Ghost `json:"ghost"`
		}{enum.GhostBlinky}
		err := json.Unmarshal(rawData, &v)
		require.Nil(t, err)
		assert.EqualValues(t, "Blinky", v.A.String())
	})
	t.Run("Unmarshal_StructField", func(t *testing.T) {
		data := `{"ghost":"Blinky"}`
		rawData := []byte(data)

		v := fooGhost{}
		err := json.Unmarshal(rawData, &v)
		require.Nil(t, err)
		assert.EqualValues(t, "Blinky", v.TestValue.String())
	})
	t.Run("Unmarshal_OmitEmptyStruct", func(t *testing.T) {
		data := `{"ghost":null}`
		rawData := []byte(data)

		v := fooGhostOmitEmpty{}
		err := json.Unmarshal(rawData, &v)
		require.NoError(t, err)
	})
	t.Run("Unmarshal_StructFieldPtr", func(t *testing.T) {
		data := `{"ghost":"Blinky"}`
		rawData := []byte(data)

		vPtr := fooGhostPtr{}
		err := json.Unmarshal(rawData, &vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, "Blinky", vPtr.TestValue.String())
	})
	t.Run("Unmarshal_StructFieldNilPtr", func(t *testing.T) {
		data := `{"ghost":null}`
		rawData := []byte(data)

		vPtr := fooGhostPtr{}
		err := json.Unmarshal(rawData, &vPtr)
		require.Nil(t, err)
		assert.Nil(t, vPtr.TestValue)
	})
}

func TestGhost_EmptyValues(t *testing.T) {
	t.Run("StructField", func(t *testing.T) {
		v := fooGhost{TestValue: enum.Ghost{}}
		data, err := json.Marshal(v)
		require.NoError(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"ghost":"%s"}`, v.TestValue.String()), string(data))

		v2 := fooGhost{}
		err = json.Unmarshal(data, &v2)
		require.NoError(t, err)
		require.True(t, v2.TestValue.IsEmpty())

		v3 := fooGhost{}
		err = json.Unmarshal(data, &v3)
		require.NoError(t, err)
		require.True(t, v3.TestValue.IsEmpty())
	})
	t.Run("PtrField", func(t *testing.T) {
		v := fooGhostPtr{TestValue: nil}
		data, err := json.Marshal(v)
		require.NoError(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"ghost":%s}`, "null"), string(data))

		v2 := fooGhost{}
		err = json.Unmarshal(data, &v2)
		require.NoError(t, err)
		require.True(t, v2.TestValue.IsEmpty())

		v3 := fooGhost{}
		err = json.Unmarshal(data, &v3)
		require.NoError(t, err)
		require.True(t, v3.TestValue.IsEmpty())
	})
}

func TestGhost_UndefinedValue(t *testing.T) {
	require.True(t, enum.Ghost{}.IsUndefined())
}

func TestGhost_ForEach(t *testing.T) {
	j := 0
	enum.EnumGhost.ForEach(func(i int, enumValue enum.Ghost) {
		j++
		value, found := enum.GhostFromValue(enumValue.String(), false)
		assert.True(t, found)
		assert.True(t, value == enumValue)

		assert.True(t, enumValue.Equals(value.String()))
		assert.True(t, enumValue.EqualsIgnoreCase(value.String()))
		assert.True(t, value.Equals(enumValue.String()))
		assert.True(t, value.EqualsIgnoreCase(enumValue.String()))
	})
	assert.EqualValues(t, enum.EnumGhost.Len(), j)
}

func TestGhost_TextCodec(t *testing.T) {
	t.Run("MarshalText_Valid", func(t *testing.T) {
		c := fooGhost{TestValue: enum.GhostBlinky}

		data, err := c.TestValue.MarshalText()
		require.Nil(t, err)
		require.True(t, len(data) > 0)
	})
	t.Run("UnmarshalText_Valid", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		err := c.TestValue.UnmarshalText([]byte("Blinky"))
		require.Nil(t, err)
		require.EqualValues(t, enum.GhostBlinky.String(), c.TestValue.String())
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		err := c.TestValue.UnmarshalText([]byte(""))
		require.NoError(t, err)
		err = c.TestValue.UnmarshalText([]byte("invalid_e6831993-894b-4957-ac84-c503d35e1a17"))
		require.NotNil(t, err)
	})
}

func TestGhost_Stringer(t *testing.T) {
	c := fooGhost{TestValue: enum.GhostBlinky}
	require.EqualValues(t, "Blinky", c.TestValue.String())
}

func TestGhost_DriverValues(t *testing.T) {
	t.Run("Scan_String", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		require.Nil(t, c.TestValue.Scan(enum.GhostBlinky.String()))
		require.EqualValues(t, enum.GhostBlinky.String(), c.TestValue.String())
	})
	t.Run("Scan_Bytes", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		require.Nil(t, c.TestValue.Scan([]byte(enum.GhostBlinky.String())))
		require.EqualValues(t, enum.GhostBlinky.String(), c.TestValue.String())
	})
	t.Run("Scan_Invalid", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		require.NotNil(t, c.TestValue.Scan(1))
	})
	t.Run("Scan_Invalid", func(t *testing.T) {
		c := fooGhost{TestValue: enum.GhostBlinky}

		v, err := c.TestValue.Value()
		assert.Nil(t, err)
		require.EqualValues(t, enum.GhostBlinky.String(), fmt.Sprintf("%v", v))
	})
}

func TestGhost_BinaryCodec(t *testing.T) {
	t.Run("MarshalBinary_Valid", func(t *testing.T) {
		c := fooGhost{TestValue: enum.GhostBlinky}

		data, err := c.TestValue.MarshalBinary()
		require.Nil(t, err)
		assert.Len(t, data, 6)
	})
	t.Run("UnmarshalBinary_Valid", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		err := c.TestValue.UnmarshalBinary([]byte("Blinky"))
		require.Nil(t, err)
		require.EqualValues(t, enum.GhostBlinky.String(), c.TestValue.String())
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		err := c.TestValue.UnmarshalBinary(nil)
		require.NoError(t, err)
		err = c.TestValue.UnmarshalBinary([]byte("invalid_e6831993-894b-4957-ac84-c503d35e1a17"))
		require.NotNil(t, err)
	})
}

func TestGhost_GobCodec(t *testing.T) {
	t.Run("MarshalBinary_Valid", func(t *testing.T) {
		c := fooGhost{TestValue: enum.GhostBlinky}

		data, err := c.TestValue.GobEncode()
		require.Nil(t, err)
		assert.Len(t, data, 6)
	})
	t.Run("UnmarshalBinary_Valid", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		err := c.TestValue.GobDecode([]byte("Blinky"))
		require.Nil(t, err)
		require.EqualValues(t, enum.GhostBlinky.String(), c.TestValue.String())
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &fooGhost{TestValue: enum.GhostBlinky}
		err := c.TestValue.GobDecode(nil)
		require.NoError(t, err)
		err = c.TestValue.GobDecode([]byte("invalid_e6831993-894b-4957-ac84-c503d35e1a17"))
		require.NotNil(t, err)
	})
}

func TestGhost_MarshalBSON(t *testing.T) {
	t.Run("MarshalBSON_Valid", func(t *testing.T) {
		c := fooGhost{TestValue: enum.GhostBlinky}

		data, err := c.TestValue.MarshalBSON()
		require.Nil(t, err)
		assert.Len(t, data, 22)
	})
}

func TestGhost_UnmarshalBSON(t *testing.T) {
	t.Run("UnmarshalBSON_Valid", func(t *testing.T) {
		c := fooGhost{TestValue: enum.GhostBlinky}

		v1 := &c
		rawData, err := bson.Marshal(v1)
		require.Nil(t, err)

		v2 := &fooGhost{}
		err = bson.Unmarshal(rawData, &v2)
		require.Nil(t, err)
		assert.EqualValues(t, "Blinky", v2.TestValue.String())
	})
}