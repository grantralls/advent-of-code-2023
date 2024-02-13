package main

import (
	"bytes"
	"testing"
)

func TestSolution1(t *testing.T) {

	type testInput struct {
		want  int
		input string
	}

	input := []testInput{
		{
			input: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			want:  142,
		},
		{
			input: "seven9rvjqdhbfour",
			want:  99,
		},
	}

	for _, test := range input {
		t.Run(test.input, func(subTest *testing.T) {
			got := Solution1(bytes.NewBufferString(test.input))

			if got != test.want {
				subTest.Errorf("Expected %d but got %d", test.want, got)
			}
		})
	}
}

func TestSolution2(t *testing.T) {

	type testInput struct {
		want  int
		input string
	}

	input := []testInput{
		{
			input: "two1nine",
			want:  29,
		},
		{
			input: "eighttwothree",
			want:  83,
		},
		{
			input: "abcone2threexyz",
			want:  13,
		},
		{
			input: "xtwone3four",
			want:  24,
		},
		{
			input: "4nineeightseven2",
			want:  42,
		},
		{
			input: "zoneight234",
			want:  14,
		},
		{
			input: "7pqrstsixteen",
			want:  76,
		},
		{
			input: "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
			want:  281,
		},
	}

	for _, test := range input {
		t.Run(test.input, func(subTest *testing.T) {
			got := Solution2(bytes.NewBufferString(test.input))

			if got != test.want {
				subTest.Errorf("Expected %d but got %d", test.want, got)
			}
		})
	}
}
