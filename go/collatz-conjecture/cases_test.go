package collatzconjecture

// Source: exercism/problem-specifications
// Commit: 25c4479 Collatz-conjecture: remove trailing space in test name (#839)
// Problem Specifications Version: 1.1.1

var testCases = []struct {
	description string
	input       int
	expectError bool
	expected    int
}{
	{
		description: "zero steps for one",
		input:       1,
		expected:    0,
	},
	{
		description: "divide if even",
		input:       16,
		expected:    4,
	},
	{
		description: "even and odd steps",
		input:       12,
		expected:    9,
	},
	{
		description: "Large number of even and odd steps",
		input:       1000000,
		expected:    152,
	},
	{
		description: "zero is an error",
		input:       0,
		expectError: true,
	},
	{
		description: "negative value is an error",
		input:       -15,
		expectError: true,
	},
}
