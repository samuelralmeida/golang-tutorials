package main

import "fmt"

/*
Declaring constraints in this way helps streamline code,
such as when a constraint is more complex.

You declare a type constraint as an interface.
The constraint allows any type implementing the interface.
For example, if you declare a type constraint interface with three methods,
then use it with a type parameter in a generic function,
type arguments used to call the function must have all of those methods.
*/
type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	// call the generic functio with type arguments
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	/*
		You can omit type arguments in calling code when the Go compiler can infer the types you want to use.
		The compiler infers type arguments from the types of function arguments.

		Note that this isn’t always possible.
		For example, if you needed to call a generic function that had no arguments,
		you would need to include the type arguments in the function call.
	*/
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

/*
Specify for the K type parameter the type constraint comparable.
Intended specifically for cases like these, the comparable constraint is predeclared in Go.
It allows any type whose values may be used as an operand of the comparison operators == and !=.
Go requires that map keys be comparable.
So declaring K as comparable is necessary so you can use K as the key in the map variable.
It also ensures that calling code uses an allowable type for map keys.
*/
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values. Usisng a type constraint.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
