package tempconv_test

import (
	"testing"

	tempconv "github.com/go-london-user-group/study-group/workspaces/jlucktay/ch2/ex2_1"
)

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
			if actual := tempconv.CToF(tC.input); tC.expected != actual {
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
			if actual := tempconv.FToC(tC.input); tC.expected != actual {
				t.Errorf("Expected '%v' but got '%v'.", tC.expected, actual)
			}
		})
	}
}
