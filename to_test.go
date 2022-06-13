package to_test

import (
	"encoding/json"
	to "github.com/rsb/from"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

type testStep struct {
	input  interface{}
	expect interface{}
	isErr  bool
}

func TestUint_Failures(t *testing.T) {
	var jMinusEight json.Number
	err := json.Unmarshal([]byte("-8"), &jMinusEight)
	require.NoError(t, err)

	var uknown struct{ foo string }
	negativeFailureMsg := to.NegativeNumberFailure.Error()
	cases := []struct {
		name string
		in   any
		msg  string
	}{
		{
			name: "int8 negative number failure",
			in:   int8(-1),
			msg:  negativeFailureMsg,
		},
		{
			name: "int16 negative number failure",
			in:   int16(-1),
			msg:  negativeFailureMsg,
		},
		{
			name: "int32 negative number failure",
			in:   int32(-1),
			msg:  negativeFailureMsg,
		},
		{
			name: "int64 negative number failure",
			in:   int64(-1),
			msg:  negativeFailureMsg,
		},
		{
			name: "float32 negative number failure",
			in:   float32(-1),
			msg:  negativeFailureMsg,
		},
		{
			name: "float64 negative number failure",
			in:   float64(-1),
			msg:  negativeFailureMsg,
		},
		{
			name: "string negative number failure",
			in:   "-1",
			msg:  negativeFailureMsg,
		},
		{
			name: "json negative number failure",
			in:   jMinusEight,
			msg:  negativeFailureMsg,
		},
		{
			name: "unknown type",
			in:   uknown,
			msg:  "unable to cast struct",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			out, err := to.Uint[uint](tt.in)
			require.Equal(t, uint(0), out)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.msg)
		})
	}

}

func createNumberTestStepsSuccess(zero, one, eight, eightNegative, eightPoint31, eightPoint31Negative interface{}) []testStep {
	var jEight, jMinusEight, jFloatEight json.Number
	_ = json.Unmarshal([]byte("8"), &jEight)
	_ = json.Unmarshal([]byte("-8"), &jMinusEight)
	_ = json.Unmarshal([]byte("8.0"), &jFloatEight)

	kind := reflect.TypeOf(zero).Kind()
	isUint := kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64

	// Some precision is lost when converting from float64 to float32.
	eightPoint3132 := eightPoint31
	eightPoint31Negative32 := eightPoint31Negative
	if kind == reflect.Float64 {
		eightPoint3132 = float64(float32(eightPoint31.(float64)))
		eightPoint31Negative32 = float64(float32(eightPoint31Negative.(float64)))
	}

	return []testStep{
		{int(8), eight, false},
		{int8(8), eight, false},
		{int16(8), eight, false},
		{int32(8), eight, false},
		{int64(8), eight, false},
		{time.Weekday(8), eight, false},
		{time.Month(8), eight, false},
		{uint(8), eight, false},
		{uint8(8), eight, false},
		{uint16(8), eight, false},
		{uint32(8), eight, false},
		{uint64(8), eight, false},
		{float32(8.31), eightPoint3132, false},
		{float64(8.31), eightPoint31, false},
		{true, one, false},
		{false, zero, false},
		{"8", eight, false},
		{nil, zero, false},
		{int(-8), eightNegative, isUint},
		{int8(-8), eightNegative, isUint},
		{int16(-8), eightNegative, isUint},
		{int32(-8), eightNegative, isUint},
		{int64(-8), eightNegative, isUint},
		{float32(-8.31), eightPoint31Negative32, isUint},
		{float64(-8.31), eightPoint31Negative, isUint},
		{"-8", eightNegative, isUint},
		{jEight, eight, false},
		{jMinusEight, eightNegative, isUint},
		{jFloatEight, eight, false},
		{"test", zero, true},
		{testing.T{}, zero, true},
	}
}
