package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSting(t *testing.T) {
	cases := map[string]struct {
		startX   int
		startY   int
		input    string
		expected [][]int
	}{
		"up2": {
			startX:   0,
			startY:   0,
			input:    "U2",
			expected: [][]int{{0, 1}, {0, 2}},
		},
		"down2": {
			startX:   0,
			startY:   0,
			input:    "D2",
			expected: [][]int{{0, -1}, {0, -2}},
		},
		"left2": {
			startX:   0,
			startY:   0,
			input:    "L2",
			expected: [][]int{{-1, 0}, {-2, 0}},
		},
		"right2": {
			startX:   0,
			startY:   0,
			input:    "R2",
			expected: [][]int{{1, 0}, {2, 0}},
		},
	}

	for n, c := range cases {
		coords := getStint(c.startX, c.startY, c.input)

		require.Equalf(t, c.expected, coords, n)
	}
}

func TestGetJourney(t *testing.T) {
	cases := map[string]struct {
		directions []string
		expected   []int
	}{
		"up2": {
			directions: []string{"U2"},
			expected:   []int{0, 2},
		},
		"down2": {
			directions: []string{"D2"},
			expected:   []int{0, -2},
		},
		"left2": {
			directions: []string{"L2"},
			expected:   []int{-2, 0},
		},
		"right2left2": {
			directions: []string{"R2", "L2"},
			expected:   []int{0, 0},
		},
		"up2down2": {
			directions: []string{"U2", "D2"},
			expected:   []int{0, 0},
		},
	}

	for n, c := range cases {
		coords, _ := getJourney(c.directions)

		require.Equalf(t, c.expected, coords, n)
	}
}
