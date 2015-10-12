package rect

import (
	"github.com/simulatedsimian/assert"
	"testing"
)

func TestMinMax(t *testing.T) {
	assert.Equal(t, min(1, 2), 1)
	assert.Equal(t, min(2, 1), 1)
	assert.Equal(t, min(1, 1), 1)
	assert.Equal(t, max(1, 2), 2)
	assert.Equal(t, max(2, 1), 2)
	assert.Equal(t, max(2, 2), 2)
}

func TestCoord(t *testing.T) {
	c := Vec{10, 20}
	c.Add(Vec{1, 1})
	assert.Equal(t, c, Vec{11, 21})
}
