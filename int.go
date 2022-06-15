package to

import (
	"golang.org/x/exp/constraints"
)

type IntData[T constraints.Integer] struct {
	item  *T
	Value string
}

func NewIntData[T constraints.Integer](v *T, sv string) IntData[T] {
	n := IntData[T]{
		item:  v,
		Value: sv,
	}

	return n
}

func (d *IntData[T]) Item() *T {
	return d.item
}
