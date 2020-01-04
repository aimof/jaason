package jaason

import (
	"testing"
)

func TestObject(t *testing.T) {
	testJSON := `{
	"object1": {
			"string0": "string0",
			"number0": 0,
			"bool0": true,
			"null0": null
		},
		"Array0": [
			1.2,
			true,
			"2"
		]
	}`

	// test Parse
	v, err := Parse([]byte(testJSON))
	if err != nil {
		t.Fatal(err)
	}

	// object
	o1 := v.Get("object1")
	if v.err != nil {
		t.Fatal()
	}
	if o1.typ != Object {
		t.Fatal()
	}

	// string
	s0, err := o1.Get("string0").String()
	if err != nil {
		t.Fatal(err)
	}
	if s0 != "string0" {
		t.Fatal(s0)
	}

	// float64
	f0, err := o1.Get("number0").Float64()
	if err != nil {
		t.Fatal(err)
	}
	if f0 != 0.0 {
		t.Fatal(f0)
	}

	// int
	i0, err := o1.Get("number0").Int()
	if err != nil {
		t.Fatal(err)
	}
	if i0 != 0 {
		t.Fatal(i0)
	}

	// bool
	b0, err := o1.Get("bool0").Bool()
	if err != nil {
		t.Fatal(err)
	}
	if !b0 {
		t.Fatal()
	}

	// null
	n0, err := o1.Get("null0").IsNull()
	if err != nil {
		t.Fatal(err)
	}
	if !n0 {
		t.Fatal()
	}

	// array
	a0 := v.Get("Array0")
	if a0.err != nil {
		t.Fatal(a0)
	}

	f1, err := a0.Get(0).Float64()
	if err != nil {
		t.Fatal(err, f1)
	}
	if f1 != 1.2 {
		t.Fatal(f1)
	}

	// index out of range
	if a0.Get(3).err == nil {
		t.Fatal()
	}

	// type
	if v.Type() != Object {
		t.Fatal()
	}
	if a0.Type() != Array {
		t.Fatal()
	}

	// When param of Get is not int nor string.
	invalidValue := v.Get('b')
	if invalidValue.err == nil {
		t.Fatal()
	}
	if v.err != nil {
		t.Fatal(v.err)
	}

	invalidValue = v.Get(3)
	if invalidValue.err == nil {
		t.Fatal()
	}
	if v.err != nil {
		t.Fatal(v.err)
	}
}
