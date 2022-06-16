package to

import (
	"encoding/json"
	"github.com/rsb/failure"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
)

type UintData[T constraints.Unsigned] struct {
	item     *T
	typeName string
}

func NewUintData[T constraints.Unsigned](v *T) UintData[T] {
	t := reflect.TypeOf(v)
	n := UintData[T]{
		item:     v,
		typeName: t.Name(),
	}

	return n
}

func (d *UintData[T]) Item() *T {
	return d.item
}

func (d *UintData[T]) Set(v string) error {
	i, err := Uint[T](v)
	if err != nil {
		return failure.Wrap(err, "Int[%v] failed", d.typeName)
	}

	d.item = &i
	return nil
}

func (d *UintData[T]) Type() string {
	return d.typeName
}

func (d *UintData[T]) String() string {
	return String(d.item)
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
		v, err := Uint[T](string(s))
		if err != nil {
			return 0, failure.ToInvalidParam(err, "Uint failed for json.Number (%v)", i)
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
