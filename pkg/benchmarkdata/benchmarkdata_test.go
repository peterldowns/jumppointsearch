package benchmarkdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/peterldowns/jumppointsearch/pkg/mapfile"
)

func TestLoadCACaveMap(t *testing.T) {
	m, err := LoadMap("ca_cave")
	require.NoError(t, err)
	assert.Equal(t, m.Type, "octile")
	assert.Equal(t, m.Height, 277)
	assert.Equal(t, m.Width, 183)
	assert.Len(t, m.Data, 277*183)
	assert.Equal(t, mapfile.TerrainPassable, m.Must(62, 48).Terrain)
	// This should be the topmost . here
	// @@@@@@@@@
	// @@@@@@@@@
	// @@@@T@@@@
	// @@TT.T@@@
	// @T...T@@T
	// T.....TTT
}

func TestLoadCACaveScenarioFile(t *testing.T) {
	scf, err := LoadScenarioFile("ca_cave")
	require.NoError(t, err)
	assert.Equal(t, "1", scf.Version)
	assert.Len(t, scf.Scenarios, 580)

	se := scf.Scenarios[0]
	assert.Equal(t, 0, se.Bucket)
	assert.Equal(t, "ca_cave.map", se.Map)
	assert.Equal(t, 183, se.MapWidth)
	assert.Equal(t, 277, se.MapHeight)
	assert.Equal(t, 122, se.StartX)
	assert.Equal(t, 208, se.StartY)
	assert.Equal(t, 125, se.GoalX)
	assert.Equal(t, 209, se.GoalY)
	assert.Equal(t, 3.41421356, se.OptimalLength)

	m, err := LoadMap(se.Map)
	require.NoError(t, err)
	assert.Equal(t, se.MapHeight, m.Height)
	assert.Equal(t, se.MapWidth, m.Width)
}

func TestLoadMapFails(t *testing.T) {
	m, err := LoadMap("does-not-exist")
	assert.Error(t, err)
	assert.Nil(t, m)
}

func TestLoadScenarioFileFails(t *testing.T) {
	sf, err := LoadScenarioFile("does-not-exist")
	assert.Error(t, err)
	assert.Nil(t, sf)
}
