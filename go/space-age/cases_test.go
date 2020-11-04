package space

// Source: exercism/problem-specifications
// Commit: 28b3dac0 space-age: restrict seconds to fit within 32-bit int range
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description string
	planet      Planet
	seconds     float64
	expected    float64
}{
	{
		description: "age on Earth",
		planet:      Planet{"Earth", 1.0},
		seconds:     1000000000,
		expected:    31.69,
	},
	{
		description: "age on Mercury",
		planet:      Planet{"Mercury", 0.2408467},
		seconds:     2134835688,
		expected:    280.88,
	},
	{
		description: "age on Venus",
		planet:      Planet{"Venus", 0.61519726},
		seconds:     189839836,
		expected:    9.78,
	},
	{
		description: "age on Mars",
		planet:      Planet{"Mars", 1.8808158},
		seconds:     2129871239,
		expected:    35.88,
	},
	{
		description: "age on Jupiter",
		planet:      Planet{"Jupiter", 11.862615},
		seconds:     901876382,
		expected:    2.41,
	},
	{
		description: "age on Saturn",
		planet:      Planet{"Saturn", 29.447498},
		seconds:     2000000000,
		expected:    2.15,
	},
	{
		description: "age on Uranus",
		planet:      Planet{"Uranus", 84.016846},
		seconds:     1210123456,
		expected:    0.46,
	},
	{
		description: "age on Neptune",
		planet:      Planet{"Neptune", 164.79132},
		seconds:     1821023456,
		expected:    0.35,
	},
}
