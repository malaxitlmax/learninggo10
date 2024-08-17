// Package ch10 provides methods for calculus operations
package ch10

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add calculate sum of two values of type Number and returns it
func Add[T Number](a, b T) T {
	return a + b
}
