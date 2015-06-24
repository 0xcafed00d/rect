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
	c := Coord{10, 20}
	c.Add(Coord{10, 10})
	testbuddy.AssertEqual(t, c, Coord{11, 11})
}
