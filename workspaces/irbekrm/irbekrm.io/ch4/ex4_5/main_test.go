package main

import (
	"reflect"
	"testing"
)

func TestRemoveAdjacentDupes(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "No duplicates found",
			input: []string{"alpha", "beta", "gamma"},
			want:  []string{"alpha", "beta", "gamma"},
		},
		{
			name:  "Duplicates found",
			input: []string{"alpha", "alpha", "beta", "gamma", "gamma", "gamma", "gamma"},
			want:  []string{"alpha", "beta", "gamma"},
		},
	}
	for _, tt := range tests {
		if got := RemoveAdjacentDupes(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("\nexpected: '%v'\ngot: '%v'\n", tt.want, got)
		}
	}

}
