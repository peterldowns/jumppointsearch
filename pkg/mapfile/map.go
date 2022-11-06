package mapfile

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Map struct {
	Type   string
	Height int
	Width  int
	Data   []Terrain
}

func (m *Map) Must(x, y int) Terrain {
	t, err := m.Get(x, y)
	if err != nil {
		panic(err)
	}
	return t
}

func (m *Map) Get(x, y int) (Terrain, error) {
	if x < 0 || x >= m.Width {
		return TerrainInvalid, fmt.Errorf("x=%d is out of [0, %d) bounds", x, m.Width)
	}
	if y < 0 || y >= m.Height {
		return TerrainInvalid, fmt.Errorf("y=%d is out of [0, %d) bounds", y, m.Height)
	}
	idx := y*m.Width + x
	return m.Data[idx], nil
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

func getMap(scanner *bufio.Scanner, width, height int) ([]Terrain, error) {
	data := make([]Terrain, 0, width*height)
	var t Terrain
	var err error
	if !scanner.Scan() {
		return nil, fmt.Errorf("scan failure")
	}
	skip := scanner.Text()
	if skip != "map" {
		return nil, fmt.Errorf("invalid map line: '%s'", skip)
	}
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			t, err = NewTerrain(string(char))
			data = append(data, t)
			if err != nil {
				break
			}
		}
	}
	return data, err
}
