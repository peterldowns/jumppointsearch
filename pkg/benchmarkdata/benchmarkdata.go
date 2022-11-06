package benchmarkdata

import (
	"embed"
	"strings"

	"github.com/peterldowns/jumppointsearch/pkg/mapfile"
)

//go:embed maps scenarios
var fs embed.FS

func LoadMap(name string) (*mapfile.Map, error) {
	path := "maps/" + name
	if !strings.HasSuffix(path, ".map") {
		path += ".map"
	}
	return mapfile.NewMapFromFS(fs, path)
}

func LoadScenarioFile(name string) (*mapfile.ScenarioFile, error) {
	path := "scenarios/" + name
	if !strings.HasSuffix(path, ".map.scen") {
		path += ".map.scen"
	}
	return mapfile.NewScenarioFileFromFS(fs, path)
}
