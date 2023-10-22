package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "part1",
			input: input,
			want:  69177,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := part1(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "part2",
			input: input,
			want:  207456,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
