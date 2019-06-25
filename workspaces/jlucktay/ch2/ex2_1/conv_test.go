package tempconv_test

import (
	"testing"

	tempconv "github.com/go-london-user-group/study-group/workspaces/jlucktay/ch2/ex2_1"
)

// floatEquals came from this gem of a gist: https://gist.github.com/cevaris/bc331cbe970b03816c6b
func floatEquals(t *testing.T, a, b tempconv.Temperature) bool {
	t.Helper()
	var EPSILON float64 = 0.00000001
	return (a.AsFloat64()-b.AsFloat64()) < EPSILON && (b.AsFloat64()-a.AsFloat64()) < EPSILON
}

func TestCtoF(t *testing.T) {
	testCases := map[string]struct {
		input    tempconv.Celsius
		expected tempconv.Fahrenheit
	}{
		"Thirty two": {
			input:    32,
			expected: 89.6,
		},
		"Two hundred and twelve": {
			input:    212,
			expected: 413.6,
		},
		"Minus forty": {
			input:    -40,
			expected: -40,
		},
	}
	for name, tC := range testCases {
		tC := tC // pin!
		t.Run(name, func(t *testing.T) {
			if actual := tempconv.CToF(tC.input); !floatEquals(t, tC.expected, actual) {
				t.Errorf("Expected '%v' but got '%v'.", tC.expected, actual)
			}
		})
	}
}

func TestCtoK(t *testing.T) {
	testCases := map[string]struct {
		input    tempconv.Celsius
		expected tempconv.Kelvin
	}{
		"Thirty two": {
			input:    32,
			expected: 305.15,
		},
		"Two hundred and twelve": {
			input:    212,
			expected: 485.15,
		},
		"Minus forty": {
			input:    -40,
			expected: 233.15,
		},
	}
	for name, tC := range testCases {
		tC := tC // pin!
		t.Run(name, func(t *testing.T) {
			if actual := tempconv.CToK(tC.input); !floatEquals(t, tC.expected, actual) {
				t.Errorf("Expected '%v' but got '%v'.", tC.expected, actual)
			}
		})
	}
}

func TestFtoC(t *testing.T) {
	testCases := map[string]struct {
		input    tempconv.Fahrenheit
		expected tempconv.Celsius
	}{
		"Thirty two": {
			input:    32,
			expected: 0,
		},
		"Two hundred and twelve": {
			input:    212,
			expected: 100,
		},
		"Minus forty": {
			input:    -40,
			expected: -40,
		},
	}
	for name, tC := range testCases {
		tC := tC // pin!
		t.Run(name, func(t *testing.T) {
			if actual := tempconv.FToC(tC.input); !floatEquals(t, tC.expected, actual) {
				t.Errorf("Expected '%v' but got '%v'.", tC.expected, actual)
			}
		})
	}
}

func TestFtoK(t *testing.T) {
	testCases := map[string]struct {
		input    tempconv.Fahrenheit
		expected tempconv.Kelvin
	}{
		"Thirty two": {
			input:    32,
			expected: 273.15,
		},
		"Two hundred and twelve": {
			input:    212,
			expected: 373.15,
		},
		"Minus forty": {
			input:    -40,
			expected: 233.15,
		},
	}
	for name, tC := range testCases {
		tC := tC // pin!
		t.Run(name, func(t *testing.T) {
			if actual := tempconv.FToK(tC.input); !floatEquals(t, tC.expected, actual) {
				t.Errorf("Expected '%v' but got '%v'.", tC.expected, actual)
			}
		})
	}
}

func TestKtoC(t *testing.T) {
	testCases := map[string]struct {
		input    tempconv.Kelvin
		expected tempconv.Celsius
	}{
		"Thirty two": {
			input:    32,
			expected: 305.15,
		},
		"Two hundred and twelve": {
			input:    212,
			expected: 485.15,
		},
		"Minus forty": {
			input:    -40,
			expected: 233.15,
		},
	}
	for name, tC := range testCases {
		tC := tC // pin!
		t.Run(name, func(t *testing.T) {
			if actual := tempconv.KToC(tC.input); !floatEquals(t, tC.expected, actual) {
				t.Errorf("Expected '%v' but got '%v'.", tC.expected, actual)
			}
		})
	}
}

func TestKtoF(t *testing.T) {
	testCases := map[string]struct {
		input    tempconv.Kelvin
		expected tempconv.Fahrenheit
	}{
		"Thirty two": {
			input:    32,
			expected: -402.07,
		},
		"Two hundred and twelve": {
			input:    212,
			expected: -78.07,
		},
		"Minus forty": {
			input:    -40,
			expected: -531.67,
		},
	}
	for name, tC := range testCases {
		tC := tC // pin!
		t.Run(name, func(t *testing.T) {
			if actual := tempconv.KToF(tC.input); !floatEquals(t, tC.expected, actual) {
				t.Errorf("Expected '%v' but got '%v'.", tC.expected, actual)
			}
		})
	}
}
