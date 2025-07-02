package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	name            string
	fizz, buzz, max int
	want            []string
}

var cases = []testCase{
	{
		name: "max flag",
		fizz: 3, buzz: 5, max: 5,
		want: []string{
			"1", "2", "Fizz", "4", "Buzz",
		},
	},
	{
		name: "defaults",
		fizz: 3, buzz: 5, max: 15,
		want: []string{
			"1", "2", "Fizz", "4", "Buzz",
			"Fizz", "7", "8", "Fizz", "Buzz",
			"11", "Fizz", "13", "14", "FizzBuzz",
		},
	},
	{
		name: "different fizz",
		fizz: 2, buzz: 5, max: 10,
		want: []string{
			"1", "Fizz", "3", "Fizz", "Buzz",
			"Fizz", "7", "Fizz", "9", "FizzBuzz",
		},
	},
	{
		name: "different buzz",
		fizz: 3, buzz: 4, max: 12,
		want: []string{
			"1", "2", "Fizz", "Buzz", "5",
			"Fizz", "7", "Buzz", "Fizz", "10",
			"11", "FizzBuzz",
		},
	},
}

func TestFizzBuzz(t *testing.T) {
	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			divisors = [2]int{tc.fizz, tc.buzz}
			max = tc.max

			got := fizzBuzz()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tc.want)
			}
		})
	}
}
