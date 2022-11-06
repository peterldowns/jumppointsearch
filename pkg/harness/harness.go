package harness

import "github.com/peterldowns/jumppointsearch/pkg/mapfile"

type Result struct {
	Path   []mapfile.Node
	Length float64
}

type Algorithm interface {
	FindPath(startX, startY, goalX, goalY int) (*Result, error)
}
