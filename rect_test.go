package rect

import (
	"github.com/simulatedsimian/testbuddy"
	"testing"
)

func TestMinMax(t *testing.T) {
	testbuddy.AssertEqual(t, min(1, 2), 1)
	testbuddy.AssertEqual(t, min(2, 1), 1)
	testbuddy.AssertEqual(t, min(1, 1), 1)
	testbuddy.AssertEqual(t, max(1, 2), 2)
	testbuddy.AssertEqual(t, max(2, 1), 2)
	testbuddy.AssertEqual(t, max(2, 2), 2)
}

func TestCoord(t *testing.T) {
	c := Vec{10, 20}
	c.Add(Vec{1, 1})
	testbuddy.AssertEqual(t, c, Vec{11, 21})
}
