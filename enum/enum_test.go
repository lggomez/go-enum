// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2020-08-02T15:36:10-03:00
// by go-enum

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fromValue(t *testing.T) {
	val, found := fromValue("Ca", true, "CountriesIso31661")
	assert.True(t, found)
	assert.False(t, val.IsUndefined())

	val, found = fromValue("cA", false, "CountriesIso31661")
	assert.False(t, found)
	assert.True(t, val.IsUndefined())

	val, found = fromValue("Ca", true, "invalid_e55bd60f-6352-4861-a0be-048acf7c3d03")
	assert.False(t, found)
	assert.True(t, val.IsUndefined())
}

func Test_stringEnumValue_Equals(t *testing.T) {
	assert.True(t, stringEnumValue{
		value: "foo",
		key:   "key",
	}.Equals("foo"))
	assert.False(t, stringEnumValue{
		value: "foo",
		key:   "key",
	}.Equals("FOO"))
}

func Test_stringEnumValue_EqualsIgnoreCase(t *testing.T) {
	assert.True(t, stringEnumValue{
		value: "foo",
		key:   "key",
	}.EqualsIgnoreCase("foo"))
	assert.True(t, stringEnumValue{
		value: "foo",
		key:   "key",
	}.EqualsIgnoreCase("FOO"))
}

func Test_stringEnumValue_IsEmpty(t *testing.T) {
	assert.True(t, stringEnumValue{}.IsEmpty())
	assert.True(t, stringEnumValue{
		value: "",
		key:   "b",
	}.IsEmpty())
	assert.False(t, stringEnumValue{
		value: "a",
		key:   "b",
	}.IsEmpty())
}

func Test_stringEnumValue_IsUndefined(t *testing.T) {
	assert.True(t, stringEnumValue{}.IsUndefined())
	assert.False(t, stringEnumValue{
		value: "a",
		key:   "b",
	}.IsUndefined())
}
