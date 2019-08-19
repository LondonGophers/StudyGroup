package main

import "unicode"

var unicodeIs map[string]struct {
	check func(rune) bool
	count uint
}

func init() {
	unicodeIs = map[string]struct {
		check func(rune) bool
		count uint
	}{
		"IsControl": {
			check: unicode.IsControl,
		},
		"IsDigit": {
			check: unicode.IsDigit,
		},
		"IsGraphic": {
			check: unicode.IsGraphic,
		},
		"IsLetter": {
			check: unicode.IsLetter,
		},
		"IsLower": {
			check: unicode.IsLower,
		},
		"IsMark": {
			check: unicode.IsMark,
		},
		"IsNumber": {
			check: unicode.IsNumber,
		},
		"IsPrint": {
			check: unicode.IsPrint,
		},
		"IsPunct": {
			check: unicode.IsPunct,
		},
		"IsSpace": {
			check: unicode.IsSpace,
		},
		"IsSymbol": {
			check: unicode.IsSymbol,
		},
		"IsTitle": {
			check: unicode.IsTitle,
		},
		"IsUpper": {
			check: unicode.IsUpper,
		},
	}
}
