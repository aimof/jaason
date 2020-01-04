package jaason

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Type is type of value in json format
type Type uint64

const (
	// Invalid format
	Invalid Type = iota
	Number
	Bool
	String
	Array
	Object
	Null
)

// Value is a parsed json value, such as Object, Array or something.
type Value struct {
	value interface{}
	typ   Type
	err   error
	fixed bool
}

// Parse json formatted string to Value.
func Parse(b []byte) (*Value, error) {
	v := new(Value)
	err := json.Unmarshal(b, &v.value)
	if err != nil {
		return nil, err
	}
	v.parse()
	return v, err
}

func (v *Value) parse() error {
	switch v.value.(type) {
	case bool:
		v.typ = Bool
	case float64:
		v.typ = Number
	case string:
		v.typ = String
	case []interface{}:
		v.typ = Array
		children := make([]*Value, 0, len(v.value.([]interface{})))
		for _, c := range v.value.([]interface{}) {
			child := new(Value)
			child.value = c
			err := child.parse()
			if err != nil {
				return err
			}
			children = append(children, child)
		}
		v.value = children
	case map[string]interface{}:
		v.typ = Object
		children := make(map[string]*Value, len(v.value.(map[string]interface{})))
		for key, c := range v.value.(map[string]interface{}) {
			child := new(Value)
			child.value = c
			err := child.parse()
			if err != nil {
				return err
			}
			children[key] = child
		}
		v.value = children
	case nil:
		v.typ = Null
	default:
		return fmt.Errorf("the type of value is invalid: %v", v.value)
	}
	return nil
}

// Get Value from object or array.
// If v is not an object or an array, it retrun nil and v.Error() != nil
func (v *Value) Get(param interface{}) *Value {
	if v.err != nil {
		return v
	}
	switch param.(type) {
	case int:
		if v.typ != Array {
			err := new(Value)
			err.err = errors.New("value is not an array")
			return err
		}
		if param.(int) < 0 || param.(int) >= len(v.value.([]*Value)) {
			err := new(Value)
			err.err = errors.New("index out of range")
			return err
		}
		children := v.value.([]*Value)[param.(int)]
		return children

	case string:
		if v.typ != Object {
			err := new(Value)
			err.err = errors.New("value is not an object")
			return err
		}
		children, ok := v.value.(map[string]*Value)[param.(string)]
		if !ok {
			err := new(Value)
			err.err = errors.New("key is not exist")
			return err
		}
		return children
	default:
		err := new(Value)
		err.err = fmt.Errorf("invalid get parameter. param=%v", param)
		return err
	}
}

// String returns value as a string if it is a string.
func (v *Value) String() (string, error) {
	if v.err != nil {
		return "", v.err
	}
	if v.typ == String {
		return v.value.(string), nil
	}
	return "", errors.New("value is not a string")
}

// Float64 returns value as float64 if it is a number.
func (v *Value) Float64() (float64, error) {
	if v.err != nil {
		return 0, v.err
	}
	if v.typ == Number {
		return v.value.(float64), nil
	}
	return 0, errors.New("value is not a number")
}

// Int returns value as int if it is a number.
func (v *Value) Int() (int, error) {
	if v.err != nil {
		return 0, v.err
	}
	if v.typ == Number {
		return int(v.value.(float64)), nil
	}
	return 0, errors.New("value is not a number")
}

// Bool returns value as bool if it is a boolean.
func (v *Value) Bool() (bool, error) {
	if v.err != nil {
		return false, v.err
	}
	if v.typ == Bool {
		return v.value.(bool), nil
	}
	return false, errors.New("value is not as boolean")
}

// IsNull returns value is null or not
func (v *Value) IsNull() (bool, error) {
	if v.err != nil {
		return false, v.err
	}
	return v.typ == Null, nil
}

// Type returns type of value
func (v *Value) Type() Type {
	return v.typ
}
