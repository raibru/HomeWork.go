package main

import (
	"testing"
)

func Test_FizzBuzz(t *testing.T) {

	// Given
	tables := []struct {
		input  int
		expect string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "17"},
	}

	for _, table := range tables {

		// When
		output := ApplyFizzBuzz(table.input)

		// Then
		if output != table.expect {
			t.Errorf("Test FittBuzz expect %s but get %s", table.expect, output)
		}
	}
}
