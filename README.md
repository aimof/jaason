# jaason: Easy to use json library for Go without defining struct.

## About

jaason is a library to use json format in Golang without defining struct.
It is simple and easy to use.
Inspired by `github.com/antonholmqist/jason`

## Usage

`go get github.com/aimof/jaason`

```go
import "gighub.com/aimof/jaason"
```

### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/aimof/jaason"
)

const j = `{
	"object1": {
		"string0": "string0",
		"number0": 0,
		"bool0": true,
		"null0": null
	},
	"array0": [
		1.2,
		true,
		"2"
	]
}`

func main() {
	rootObject, err := jaason.Parse([]byte(j))
	// handle error

	s, err := rootObject.Get("object1").Get("string0").String()
	// handle error
	fmt.Printf("%s\n", s) // string0

	f, err := rootObject.Get("array0").Get(0).Float64()
	// handle error
	fmt.Printf("%.1f\n", f) // 1.2
}
```

## Usage

### tyep Value

```go
type Value struct {
    // contains filtered or unexported fields
}
```

Value is a parsed JSON value.
It may be an object, array, boolean, number, string or null.

### Parse

```go
func Parse(b []byte) (*Value, error)
```

Parse JSON formatted bytes to *Value.

### Get

```go 
func (*Value) Get(interface{}) *Value
```

Get function is like JS code below.

```js
object["key"]
array[0]
```

When Get from Value with key or index, it returns the value.

The parameter must be a string or a int.

**Get doesn't return error.**
It has errors in value.
IT returns errors when calls functions below.

### type functions

Get returns type *jaason.Value so you must use these functions.

```go
func (*Value) Bool() (bool, error)
func (*Value) Int() (int, error)
func (*Value) String() (string, error)
func (*Value) IsNull() (bool, error)
```

## ToDo

- [ ] `func (*Value) Err() error`.
- [ ] Write more test cases.
- [ ] Write more README.md.
- [ ] Define errors.

## LICNSE

MIT