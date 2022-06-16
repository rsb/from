package to

import (
	"encoding/json"
	"github.com/rsb/failure"
	"golang.org/x/exp/constraints"
	"strconv"
)

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
