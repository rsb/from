// Package to is responsible for casting between go types
package to

import (
	"encoding/json"
	"fmt"
	"github.com/rsb/failure"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

type Number interface {
	int | int8
}

var (
	NegativeNumberFailure = failure.InvalidParam("negative numbers permitted")
)

func Int[T constraints.Signed](i any) (T, error) {
	i = indirect(i)

	v, ok := integer(i)
	if ok {
		return T(v), nil
	}

	switch s := i.(type) {
	case int8:
		return T(s), nil
	case int16:
		return T(s), nil
	case int32:
		return T(s), nil
	case int64:
		return T(s), nil
	case uint:
		return T(s), nil
	case uint8:
		return T(s), nil
	case uint16:
		return T(s), nil
	case uint32:
		return T(s), nil
	case uint64:
		return T(s), nil
	case float32:
		return T(s), nil
	case float64:
		return T(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0, failure.ToInvalidParam(err, "unable to cast %#v of type %T to int64", i, i)
		}
		return T(v), nil
	case json.Number:
		v, err := Int(string(s))
		if err != nil {
			return 0, failure.ToInvalidParam(err, "Int failed for json.Number (%v)", i)
		}
		return v, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, failure.InvalidParam("unable to cast %#v of type %T to int", i, i)
	}
}

func Uint[T constraints.Unsigned](i any) (T, error) {
	i = indirect(i)

	v, ok := integer(i)
	if ok {
		if v < 0 {
			return 0, NegativeNumberFailure
		}

		return T(v), nil
	}

	switch s := i.(type) {
	case int8:
		if s < 0 {
			return 0, NegativeNumberFailure
		}
		return T(s), nil
	case int16:
		if s < 0 {
			return 0, NegativeNumberFailure
		}
		return T(s), nil
	case int32:
		if s < 0 {
			return 0, NegativeNumberFailure
		}
		return T(s), nil
	case int64:
		if s < 0 {
			return 0, NegativeNumberFailure
		}
		return T(s), nil
	case uint:
		return T(s), nil
	case uint8:
		return T(s), nil
	case uint16:
		return T(s), nil
	case uint32:
		return T(s), nil
	case uint64:
		return T(s), nil
	case float32:
		if s < 0 {
			return 0, NegativeNumberFailure
		}
		return T(s), nil
	case float64:
		if s < 0 {
			return 0, NegativeNumberFailure
		}
		return T(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0, failure.ToInvalidParam(err, "unable to cast %#v of type %T to uint", i, i)
		}
		if v < 0 {
			return 0, NegativeNumberFailure
		}
		return T(v), nil
	case json.Number:
		v, err := Uint(string(s))
		if err != nil {
			return 0, failure.ToInvalidParam(err, "Int failed for json.Number (%v)", i)
		}
		return v, nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, failure.InvalidParam("unable to cast %#v of type %T to int", i, i)
	}
}

func Float[T constraints.Float](i any) (T, error) {
	i = indirect(i)

	v, ok := integer(i)
	if ok {
		return T(v), nil
	}

	switch s := i.(type) {
	case int8:
		return T(s), nil
	case int16:
		return T(s), nil
	case int32:
		return T(s), nil
	case int64:
		return T(s), nil
	case uint:
		return T(s), nil
	case uint8:
		return T(s), nil
	case uint16:
		return T(s), nil
	case uint32:
		return T(s), nil
	case uint64:
		return T(s), nil
	case float32:
		return T(s), nil
	case float64:
		return T(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, failure.ToInvalidParam(err, "unable to cast %#v of type %T to int64", i, i)
		}
		return T(v), nil
	case json.Number:
		v, err := s.Float64()
		if err != nil {
			return 0, failure.ToInvalidParam(err, "Int failed for json.Number (%v)", i)
		}
		return T(v), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, failure.InvalidParam("unable to cast %#v of type %T to int", i, i)
	}
}

// integer returns the int value of v if v or v's underlying type
// is an int.
// Note that this will return false for int64 etc. types.
func integer(v interface{}) (int, bool) {
	switch v := v.(type) {
	case int:
		return v, true
	case time.Weekday:
		return int(v), true
	case time.Month:
		return int(v), true
	default:
		return 0, false
	}
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after de-referencing as many times
// as necessary to reach the base type (or nil).
func indirect(a any) any {
	if a == nil {
		return nil
	}

	if t := reflect.TypeOf(a); t.Kind() != reflect.Pointer {
		// Avoid creating a `reflect.Value` if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after de-referencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
