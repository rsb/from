// Package to is responsible for casting between go types
package to

import (
	"fmt"
	"github.com/rsb/failure"
	"reflect"
)

var (
	NegativeNumberFailure = failure.InvalidParam("negative numbers are not permitted")
)

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
func indirectToStringerOrError(a any) interface{} {
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
