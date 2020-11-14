package listops

import (
	"testing"
)

var foldTestCases = []struct {
	name     string
	property string
	fn       func(a, b int) int
	initial  int
	list     []int
	want     int
}{
	{
		name:     "empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     0,
		list:     []int{2, 5},
	},
	{
		name:     "empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     2,
		list:     []int{2, 5},
	},
}

func TestFold(t *testing.T) {
	var got int
	for _, tt := range foldTestCases {
		list := MakeList(tt.list...)
		if tt.property == "foldr" {
			got = list.Foldr(tt.fn, tt.initial)
		} else {
			got = list.Foldl(tt.fn, tt.initial)
		}
		if got != tt.want {
			t.Fatalf("FAIL: %s: %q -- expected: %d, actual: %d", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

var filterTestCases = []struct {
	name     string
	property string
	fn       func(int) bool
	list     []int
	want     []int
}{
	{
		name:     "empty list",
		property: "filter",
		fn:       func(n int) bool { return n%2 == 1 },
		list:     []int{},
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "filter",
		fn:       func(n int) bool { return n%2 == 1 },
		list:     []int{1, 2, 3, 4, 5},
		want:     []int{1, 3, 5},
	},
}

func TestFilterMethod(t *testing.T) {
	for _, tt := range filterTestCases {
		in := MakeList(tt.list...)
		got := in.Filter(tt.fn)
		want := MakeList(tt.want...)
		if !Equal(want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

var lengthTestCases = []struct {
	name     string
	property string
	list     []int
	want     int
}{
	{
		name:     "empty list",
		property: "length",
		list:     []int{},
		want:     0,
	},
	{
		name:     "non-empty list",
		property: "length",
		list:     []int{1, 2, 3, 4},
		want:     4,
	},
}

func TestLengthMethod(t *testing.T) {
	for _, tt := range lengthTestCases {
		got := MakeList(tt.list...).Length()
		if tt.want != got {
			t.Fatalf("FAIL: %s: %q -- expected: %d, actual: %d", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

var mapTestCases = []struct {
	name     string
	property string
	list     []int
	fn       func(int) int
	want     []int
}{
	{
		name:     "empty list",
		property: "map",
		list:     []int{},
		fn:       func(x int) int { return x + 1 },
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "map",
		list:     []int{1, 3, 5, 7},
		fn:       func(x int) int { return x + 1 },
		want:     []int{2, 4, 6, 8},
	},
}

func TestMapMethod(t *testing.T) {
	for _, tt := range mapTestCases {
		got := MakeList(tt.list...).Map(tt.fn)
		want := MakeList(tt.want...)
		if !Equal(want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}

	}
}

var reverseTestCases = []struct {
	name     string
	property string
	list     []int
	want     []int
}{
	{
		name:     "empty list",
		property: "reverse",
		list:     []int{},
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "reverse",
		list:     []int{1, 3, 5, 7},
		want:     []int{7, 5, 3, 1},
	},
}

func TestReverseMethod(t *testing.T) {
	for _, tt := range reverseTestCases {
		got := MakeList(tt.list...).Reverse()
		want := MakeList(tt.want...)
		if !Equal(want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

var appendTestCases = []struct {
	name       string
	property   string
	list       []int
	appendThis []int
	want       []int
}{
	{
		name:       "empty list",
		property:   "append",
		list:       []int{},
		appendThis: []int{},
		want:       []int{},
	},
	{
		name:       "empty list to list",
		property:   "append",
		list:       []int{},
		appendThis: []int{1, 2, 3, 4},
		want:       []int{1, 2, 3, 4},
	},
	{
		name:       "non-empty lists",
		property:   "append",
		list:       []int{1, 2},
		appendThis: []int{2, 3, 4, 5},
		want:       []int{1, 2, 2, 3, 4, 5},
	},
}

func TestAppendMethod(t *testing.T) {
	for _, tt := range appendTestCases {
		got := MakeList(tt.list...).Append(MakeList(tt.appendThis...))
		want := MakeList(tt.want...)
		if !Equal(want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

var concatTestCases = []struct {
	name     string
	property string
	list     []int
	args     [][]int
	want     []int
}{
	{
		name:     "empty list",
		property: "concat",
		list:     []int{},
		args:     [][]int{},
		want:     []int{},
	},
	{
		name:     "list of lists",
		property: "concat",
		list:     []int{1, 2},
		args:     [][]int{[]int{3}, []int{}, []int{4, 5, 6}},
		want:     []int{1, 2, 3, 4, 5, 6},
	},
}

func TestConcatMethod(t *testing.T) {
	for _, tt := range concatTestCases {
		var args []*IntList
		for _, a := range tt.args {
			args = append(args, MakeList(a...))
		}
		got := MakeList(tt.list...).Concat(args...)
		want := MakeList(tt.want...)
		if !Equal(want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}
