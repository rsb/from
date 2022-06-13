// Package to is responsible for casting between go types
package to

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/rsb/failure"
)

func Int(i any) (int, error) {
	i = indirect(i)

	v, ok := integer(i)
	if ok {
		return v, nil
	}

	switch s := i.(type) {
	case int8:
		return int(s), nil
	case int16:
		return int(s), nil
	case int32:
		return int(s), nil
	case int64:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint8:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint64:
		return int(s), nil
	case float32:
		return int(s), nil
	case float64:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0, failure.ToInvalidParam(err, "unable to cast %#v of type %T to int64", i, i)
		}
		return int(v), nil
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
// indirect returns the value, after dereferencing as many times
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
