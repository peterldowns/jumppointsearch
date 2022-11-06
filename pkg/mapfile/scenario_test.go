package mapfile

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewScenario(t *testing.T) {
	raw := strings.TrimSpace(`
1	example.map	100	200	1	2	99	199	99.9999
`)
	se, err := NewScenario(raw)
	require.NoError(t, err)

	expected := Scenario{
		Bucket:        1,
		Map:           "example.map",
		MapWidth:      100,
		MapHeight:     200,
		StartX:        1,
		StartY:        2,
		GoalX:         99,
		GoalY:         199,
		OptimalLength: 99.9999,
	}
	require.Equal(t, expected, *se)
}
