package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/peterldowns/jumppointsearch/pkg/algorithms"
	"github.com/peterldowns/jumppointsearch/pkg/benchmarkdata"
	"github.com/peterldowns/jumppointsearch/pkg/mapfile"
)

func TestSimpleDFS(t *testing.T) {
	m, err := mapfile.NewMap(strings.TrimSpace(`
type octile
height 8
width 9
map
@@@@@@@@@
@.@.....@
@.@.@@@@@
@.@.....@
@.@.@TTT@
@.@.@TTT@
@...@TTT@
@@@@@@@@@
`))
	require.NoError(t, err)
	dfs := algorithms.NewDFS(*m)
	result, err := dfs.FindPath(1, 1, 7, 3)
	require.NoError(t, err)

	rendered := m.Render(result.Path)
	expected := strings.TrimSpace(`
@@@@@@@@@
@'@.....@
@'@.@@@@@
@'@'''''@
@'@'@TTT@
@'@'@TTT@
@'''@TTT@
@@@@@@@@@
`)
	if expected != rendered {
		fmt.Println("expected:")
		fmt.Println("")
		fmt.Println(expected)
		fmt.Println("")
		fmt.Println("received:")
		fmt.Println("")
		fmt.Println(rendered)
		fmt.Println("")
		t.Errorf("Did not find expected path")
	}
}

func TestAllScenarios_ca_cave(t *testing.T) {
	scf, err := benchmarkdata.LoadScenarioFile("ca_cave")
	require.NoError(t, err)
	m, err := benchmarkdata.LoadMap(scf.Scenarios[0].Map)
	require.NoError(t, err)
	dfs := algorithms.NewDFS(*m)

	for i, se := range scf.Scenarios {
		_, err := dfs.FindPath(se.StartX, se.StartY, se.GoalX, se.GoalY)
		assert.NoError(t, err, fmt.Sprintf("scenario %d failed", i))
	}
}
