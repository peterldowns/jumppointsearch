package mapfile

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Node struct {
	X       int
	Y       int
	Terrain Terrain
}

type Map struct {
	Type   string
	Height int
	Width  int
	Data   []Node
}

func (m *Map) Must(x, y int) Node {
	n, err := m.Get(x, y)
	if err != nil {
		panic(err)
	}
	return n
}

func (m *Map) Get(x, y int) (Node, error) {
	if x < 0 || x >= m.Width {
		return Node{}, fmt.Errorf("x=%d is out of [0, %d) bounds", x, m.Width)
	}
	if y < 0 || y >= m.Height {
		return Node{}, fmt.Errorf("y=%d is out of [0, %d) bounds", y, m.Height)
	}
	idx := y*m.Width + x
	return m.Data[idx], nil
}

func (m *Map) Render(path []Node) string {
	var b bytes.Buffer

	inPath := map[Node]struct{}{}
	for _, node := range path {
		inPath[node] = struct{}{}
	}
	for i, node := range m.Data {
		if i != 0 && i%m.Width == 0 {
			b.WriteRune('\n')
		}
		if _, ok := inPath[node]; ok {
			b.WriteRune('\'')
		} else {
			b.WriteString(string(node.Terrain))
		}
	}
	return b.String()
}

func (m *Map) Neighbors(n Node) []Node {
	var valid []Node
	for _, candidate := range []Node{
		{n.X - 1, n.Y, TerrainInvalid},
		{n.X + 1, n.Y, TerrainInvalid},
		{n.X, n.Y - 1, TerrainInvalid},
		{n.X, n.Y + 1, TerrainInvalid},
		// Diagonals not yet implemented
		// {n.X - 1, n.Y - 1, TerrainInvalid},
		// {n.X - 1, n.Y + 1, TerrainInvalid},
		// {n.X + 1, n.Y + 1, TerrainInvalid},
		// {n.X + 1, n.Y - 1, TerrainInvalid},
	} {
		tmp, err := m.Get(candidate.X, candidate.Y)
		if err == nil {
			valid = append(valid, tmp)
		}
	}
	return valid
}

func NewMapFromFS(fs embed.FS, path string) (*Map, error) {
	file, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return NewMapFromReader(file)
}

func NewMapFromReader(r io.Reader) (*Map, error) {
	scanner := bufio.NewScanner(r)
	return newMap(scanner)
}

func NewMap(raw string) (*Map, error) {
	scanner := bufio.NewScanner(strings.NewReader(raw))
	return newMap(scanner)
}

func newMap(scanner *bufio.Scanner) (*Map, error) {
	m := &Map{}
	var err error
	m.Type, err = getType(scanner)
	if err != nil {
		return m, err
	}
	m.Height, err = getHeight(scanner)
	if err != nil {
		return m, err
	}
	m.Width, err = getWidth(scanner)
	if err != nil {
		return m, err
	}
	m.Data, err = getMap(scanner, m.Width, m.Height)
	if err != nil {
		return m, err
	}
	return m, nil
}

func getType(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", fmt.Errorf("scan failure")
	}
	text := scanner.Text()
	data := strings.Fields(text)
	if len(data) != 2 {
		return "", fmt.Errorf("invalid type line: '%s'", text)
	}
	return data[1], nil
}

func getHeight(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		return -1, fmt.Errorf("scan failure")
	}
	text := scanner.Text()
	data := strings.Fields(text)
	if len(data) != 2 || data[0] != "height" {
		return -1, fmt.Errorf("invalid height line: '%s'", text)
	}
	i, err := strconv.ParseInt(data[1], 10, 0)
	if err != nil {
		return -1, fmt.Errorf("invalid height value: '%s'", data[1])
	}
	return int(i), nil
}

func getWidth(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		return -1, fmt.Errorf("scan failure")
	}
	text := scanner.Text()
	data := strings.Fields(text)
	if len(data) != 2 || data[0] != "width" {
		return -1, fmt.Errorf("invalid width line: '%s'", text)
	}
	i, err := strconv.ParseInt(data[1], 10, 0)
	if err != nil {
		return -1, fmt.Errorf("invalid width value: '%s'", data[1])
	}
	return int(i), nil
}

func getMap(scanner *bufio.Scanner, width, height int) ([]Node, error) {
	data := make([]Node, 0, width*height)
	var t Terrain
	var err error
	if !scanner.Scan() {
		return nil, fmt.Errorf("scan failure")
	}
	skip := scanner.Text()
	if skip != "map" {
		return nil, fmt.Errorf("invalid map line: '%s'", skip)
	}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		x := 0
		for _, char := range line {
			t, err = NewTerrain(string(char))
			if err != nil {
				break
			}
			data = append(data, Node{X: x, Y: y, Terrain: t})
			x++
		}
		y++
	}
	return data, err
}
