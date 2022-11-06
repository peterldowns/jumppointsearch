package mapfile

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMap(t *testing.T) {
	raw := strings.TrimSpace(`
type octile
height 3
width 5
map
O@@@O
TWS.G
O@@@O
`)
	m, err := NewMap(raw)
	require.NoError(t, err)

	assert.Equal(t, "octile", m.Type)
	assert.Equal(t, 3, m.Height)
	assert.Equal(t, 5, m.Width)
	assert.Len(t, m.Data, 15)

	for y, values := range [][]Terrain{
		{
			TerrainOutOfBounds, // canonicalized from "O"
			TerrainOutOfBounds,
			TerrainOutOfBounds,
			TerrainOutOfBounds,
			TerrainOutOfBounds, // canonicalized from "O"
		},
		{
			TerrainTrees,
			TerrainWater,
			TerrainSwamp,
			TerrainPassable,
			TerrainPassable,
		},
		{
			TerrainOutOfBounds, // canonicalized from "O"
			TerrainOutOfBounds,
			TerrainOutOfBounds,
			TerrainOutOfBounds,
			TerrainOutOfBounds, // canonicalized from "O"
		},
	} {
		for x, expected := range values {
			actual, err := m.Get(x, y)
			assert.NoError(t, err)
			assert.Equal(t, expected, actual.Terrain, fmt.Sprintf("x=%d y=%d", x, y))
		}
	}

	// 5 is out of bounds, max value 4
	_, err = m.Get(5, 2)
	assert.ErrorContains(t, err, "[0, 5)")
	// 3 is out of bounds, max value 2
	_, err = m.Get(4, 3)
	assert.ErrorContains(t, err, "[0, 3)")
}
