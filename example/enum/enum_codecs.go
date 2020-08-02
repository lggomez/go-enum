// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2020-08-02T15:35:01-03:00
// by go-enum

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// codecs: marshal/unmarshal methods for several native interfaces:
// 	- Stringer
// 	- json.Marshaler, json.Unmarshaler
// 	- text.Marshaler, text.Unmarshaler
// 	- bson.Marshaler, bson.Unmarshaler
// 	- json.Marshaler, json.Unmarshaler
// 	- gob.GobEncoder, gob.GobDecoder
// 	- driver.Valuer, sql.Scanner

const (
	JSONNull string = "null"
)

// Validate and assign value to unmarshal into target enum instance.
//
// Since we know the specific type by the key injected in e.key, we
// check against the enum index to validate the incoming value
func (e *stringEnumValue) validateValueByKey(value string) error {
	if _, ok := enumIndex[e.key][value]; !ok && (value != "") {
		return fmt.Errorf("stringEnumValue: value '%v' is not allowed", value)
	}
	return nil
}

// Stringer implementation
func (e stringEnumValue) String() string { return e.value }

// MarshalJSON returns the stringEnumValue value as JSON
func (e stringEnumValue) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(e.value)
	return data, err
}

// UnmarshalJSON sets the stringEnumValue value from JSON
func (e *stringEnumValue) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if err := e.validateValueByKey(value); err != nil {
		return err
	}

	e.value = value
	return nil
}

// UnmarshalText parses a text representation into a date types
func (e *stringEnumValue) UnmarshalText(text []byte) error {
	value := string(text)

	if err := e.validateValueByKey(value); err != nil {
		return err
	}

	e.value = value
	return nil
}

// MarshalText serializes this date types to string
func (e stringEnumValue) MarshalText() ([]byte, error) {
	data := []byte(e.String())
	return data, nil
}

// Scan scans a stringEnumValue value from database driver types.
func (e *stringEnumValue) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case []byte:
		return e.UnmarshalText(v)
	case string:
		return e.UnmarshalText([]byte(v))
	default:
		return fmt.Errorf("cannot sql.Scan() enum.stringEnumValue from: %#v", v)
	}
}

// Value converts stringEnumValue to a primitive value ready to written to a database.
func (e stringEnumValue) Value() (driver.Value, error) {
	return driver.Value(e.String()), nil
}

// MarshalBSON implements the bson.Marshaler interface.
func (e stringEnumValue) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bson.M{"data": e.String()})
}

// UnmarshalBSON implements the bson.Unmarshaler interface.
func (e *stringEnumValue) UnmarshalBSON(data []byte) error {
	var m bson.M
	if err := bson.Unmarshal(data, &m); err != nil {
		return err
	}

	if data, ok := m["data"].(string); ok {
		if err := e.validateValueByKey(data); err != nil {
			return err
		}
		e.value = data
		return nil
	}

	return errors.New("couldn't unmarshal bson bytes string as enum.stringEnumValue")
}

// GobEncode implements the gob.GobEncoder interface.
func (e stringEnumValue) GobEncode() ([]byte, error) {
	return e.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface.
func (e *stringEnumValue) GobDecode(data []byte) error {
	return e.UnmarshalBinary(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (e stringEnumValue) MarshalBinary() ([]byte, error) {
	return []byte(e.value), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (e *stringEnumValue) UnmarshalBinary(data []byte) error {
	value := string(data)
	if err := e.validateValueByKey(value); err != nil {
		return err
	}
	e.value = value
	return nil
}
