package to_test

import (
	"encoding/json"
	to "github.com/rsb/from"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUint_SingleIndirect(t *testing.T) {
	t.Parallel()
	var foo int32

	foo = int32(100)

	out, err := to.Uint[uint](&foo)
	require.NoError(t, err)
	require.Equal(t, uint(100), out)
}

func TestUint_FourLevelsOfIndirect(t *testing.T) {
	t.Parallel()
	var foo int32

	foo = int32(100)
	bar := &foo
	baz := &bar
	biz := &baz
	out, err := to.Uint[uint](&biz)
	require.NoError(t, err)
	require.Equal(t, uint(100), out)
}

func TestUint_Success(t *testing.T) {
	t.Parallel()
	var jEight json.Number

	err := json.Unmarshal([]byte("8"), &jEight)
	require.NoError(t, err)

	cases := []struct {
		name     string
		in       any
		expected uint
	}{
		{
			name:     "int",
			in:       8,
			expected: uint(8),
		},
		{
			name:     "int8",
			in:       int8(8),
			expected: uint(8),
		},
		{
			name:     "int16",
			in:       int16(8),
			expected: uint(8),
		},
		{
			name:     "int32",
			in:       int32(8),
			expected: uint(8),
		},
		{
			name:     "int64",
			in:       int64(8),
			expected: uint(8),
		},
		{
			name:     "uint",
			in:       uint(8),
			expected: uint(8),
		},
		{
			name:     "uint8",
			in:       uint8(8),
			expected: uint(8),
		},
		{
			name:     "uint16",
			in:       uint16(8),
			expected: uint(8),
		},
		{
			name:     "uint32",
			in:       uint32(8),
			expected: uint(8),
		},
		{
			name:     "uint64",
			in:       uint64(8),
			expected: uint(8),
		},
		{
			name:     "float32",
			in:       float32(8.8),
			expected: uint(8),
		},
		{
			name:     "float64",
			in:       8.8,
			expected: uint(8),
		},
		{
			name:     "json.Number 8",
			in:       jEight,
			expected: uint(8),
		},
		{
			name:     "bool true",
			in:       true,
			expected: uint(1),
		},
		{
			name:     "bool false",
			in:       false,
			expected: uint(0),
		},
		{
			name:     "nil",
			in:       nil,
			expected: uint(0),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			out, err := to.Uint[uint](tt.in)
			require.NoError(t, err)
			require.Equal(t, tt.expected, out)
		})
	}
}

func TestUint_Failures(t *testing.T) {
	t.Parallel()

	var jMinusEight, jEightDotEight json.Number
	err := json.Unmarshal([]byte("-8"), &jMinusEight)
	require.NoError(t, err)
	err = json.Unmarshal([]byte("8.8"), &jEightDotEight)
	require.NoError(t, err)

	var uKnown struct{ foo string }
	negativeFailureMsg := to.NegativeNumberFailure.Error()
	cases := []struct {
		name string
		in   any
		msg  string
	}{
		{
			name: "int negative number failure",
			in:   -1,
			msg:  negativeFailureMsg,
		},
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
			name: "json float",
			in:   jEightDotEight,
			msg:  "Uint failed for json.Number (8.8)",
		},
		{
			name: "unknown type",
			in:   uKnown,
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
