package mapfile

import (
	"fmt"
)

type Terrain string

const (
	TerrainPassable    Terrain = "."
	TerrainOutOfBounds Terrain = "@"
	TerrainTrees       Terrain = "T"
	TerrainSwamp       Terrain = "S"
	TerrainWater       Terrain = "W"
	TerrainInvalid     Terrain = "X"
	altPassable        Terrain = "G"
	altOutOfBounds     Terrain = "O"
)

var canonical = map[Terrain]Terrain{ //nolint: gochecknoglobals
	TerrainPassable:    TerrainPassable,
	TerrainOutOfBounds: TerrainOutOfBounds,
	TerrainTrees:       TerrainTrees,
	TerrainSwamp:       TerrainSwamp,
	TerrainWater:       TerrainWater,
	TerrainInvalid:     TerrainInvalid,
	altPassable:        TerrainPassable,
	altOutOfBounds:     TerrainOutOfBounds,
}

func NewTerrain(char string) (Terrain, error) {
	t, ok := canonical[Terrain(char)]
	if !ok {
		return TerrainInvalid, fmt.Errorf("invalid terrain character: '%s'", char)
	}
	return t, nil
}
