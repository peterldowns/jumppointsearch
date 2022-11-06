package algorithms

import (
	"fmt"

	"github.com/peterldowns/jumppointsearch/pkg/harness"
	"github.com/peterldowns/jumppointsearch/pkg/mapfile"
)

type DFS struct {
	Map mapfile.Map
}

func NewDFS(m mapfile.Map) *DFS {
	dfs := &DFS{Map: m}
	return dfs
}

func (dfs *DFS) FindPath(startX, startY, goalX, goalY int) (*harness.Result, error) {
	parents := map[mapfile.Node]mapfile.Node{}
	lengths := map[mapfile.Node]int{}
	goal, err := dfs.Map.Get(goalX, goalY)
	if err != nil {
		return nil, fmt.Errorf("could not get goal node: %w", err)
	}
	start, err := dfs.Map.Get(startX, startY)
	if err != nil {
		return nil, fmt.Errorf("could not get start node: %w", err)
	}

	// TerrainInvalid is used as the stop condition when backtracing from goal
	// to construct the path.
	parents[start] = mapfile.Node{X: -1, Y: -1, Terrain: mapfile.TerrainInvalid}
	lengths[start] = 1

	q := &[]mapfile.Node{start}
	for len(*q) > 0 {
		currentNode := pop(q)
		if currentNode == goal {
			break
		}
		currentLength := lengths[currentNode]
		children := neighbors(&dfs.Map, currentNode)
		for _, child := range children {
			if _, visited := parents[child]; visited {
				continue
			}
			if child.Terrain == mapfile.TerrainPassable {
				push(q, child)
				parents[child] = currentNode
				lengths[child] = currentLength + 1
			}
			if child == goal {
				break
			}
		}
	}

	lengthOfPath, reachedGoal := lengths[goal]
	if !reachedGoal {
		return nil, fmt.Errorf("search did not reach goal")
	}
	result := &harness.Result{
		Path: make([]mapfile.Node, lengthOfPath),
		// Not doing diagonals or distances yet.
		Length: float64(lengthOfPath),
	}
	for i, n := 1, goal; n.Terrain != mapfile.TerrainInvalid; i++ {
		result.Path[lengthOfPath-i] = n
		n = parents[n]
	}
	return result, nil
}

// push the node `n` to the stack
func push(q *[]mapfile.Node, n mapfile.Node) {
	*q = append(*q, n)
}

// pop the most-recently-added node `n` from the stack
func pop(q *[]mapfile.Node) mapfile.Node {
	var n mapfile.Node
	count := len(*q)
	n, *q = (*q)[count-1], (*q)[0:count-1]
	return n
}

func neighbors(m *mapfile.Map, n mapfile.Node) []mapfile.Node {
	var valid []mapfile.Node
	if node, err := m.Get(n.X, n.Y-1); err == nil {
		valid = append(valid, node)
	}
	if node, err := m.Get(n.X-1, n.Y); err == nil {
		valid = append(valid, node)
	}
	if node, err := m.Get(n.X, n.Y+1); err == nil {
		valid = append(valid, node)
	}
	if node, err := m.Get(n.X+1, n.Y); err == nil {
		valid = append(valid, node)
	}
	return valid
}
