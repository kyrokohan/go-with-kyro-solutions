package main

import (
	"testing"
)

type testCase struct {
	name    string
	noColor bool
	want    string
}

var cases = []testCase{
	{
		// default case
		name: "World", noColor: false,
		want: "Hello, \033[36mWorld\033[0m!",
	},
	{
		// different name test
		name: "Kyro", noColor: false,
		want: "Hello, \033[36mKyro\033[0m!",
	},
	{
		// no color test
		name: "Alice", noColor: true,
		want: "Hello, Alice!",
	},
}

func TestHelloCLI(t *testing.T) {
	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			if tc.noColor {
				t.Setenv("NO_COLOR", "true")
			} else {
				t.Setenv("NO_COLOR", "")
			}

			got := greet(tc.name, tc.noColor)

			if got != tc.want {
				t.Errorf("greet(...) = %q, want %q", got, tc.want)
			}

		})

	}
}
