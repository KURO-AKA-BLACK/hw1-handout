package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	type Test struct {
		a, b, c  int
		expected triangleType
	}

	var tests = []Test{
		{1001, 5, 6, UnknownTriangle},
		// TODO add more tests for 100% test coverage
		{999, 2001, 6, UnknownTriangle},
		{5, 6, 3001, UnknownTriangle},
		{-1, 6, 5, UnknownTriangle},
		{1, -1, 9, UnknownTriangle},
		{5, 6, -1, UnknownTriangle},
		{1, 1, 4, InvalidTriangle},
		{3, 4, 5, RightTriangle},
		{4, 5, 6, AcuteTriangle},
		{2, 2, 3, ObtuseTriangle},
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
