package reflection

import (
	// "fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
// fmt.Printf("Type: %v, Kind: %v, Value: %v\n", val.Type(), val.Kind(), val)
	field := val.Field(0)
// fmt.Printf("Type: %v, Kind: %v, Value: %v\n", field.Type(), field.Kind(), field)
	fn(field.String())
}