package main

// Convertable types take a float64 and return a string representation of
// the value in various alternate units.
// THOUGHT: would it be better to return structured data?
// THOUGHT: if this returned a custom type, where should it be defined?
type Convertable interface {
	Convert(float64) (string, error)
	Name() string
}

var conversions = []Convertable{}

func register(c Convertable) {
	conversions = append(conversions, c)
}
