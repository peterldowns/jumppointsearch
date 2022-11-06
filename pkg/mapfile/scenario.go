package mapfile

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type ScenarioFile struct {
	Version   string
	Scenarios []Scenario
}

func NewScenarioFileFromFS(fs embed.FS, path string) (*ScenarioFile, error) {
	file, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return NewScenarioFileFromReader(file)
}

func NewScenarioFileFromReader(r io.Reader) (*ScenarioFile, error) {
	scanner := bufio.NewScanner(r)
	return newScenarioFile(scanner)
}

func newScenarioFile(scanner *bufio.Scanner) (*ScenarioFile, error) {
	sf := &ScenarioFile{}
	var err error
	sf.Version, err = getVersion(scanner)
	if err != nil {
		return nil, err
	}
	sf.Scenarios, err = getScenarios(scanner)
	if err != nil {
		return nil, err
	}
	return sf, nil
}

func getVersion(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", fmt.Errorf("scan failure")
	}
	text := scanner.Text()
	data := strings.Fields(text)
	if len(data) != 2 {
		return "", fmt.Errorf("invalid version line: '%s'", text)
	}
	return data[1], nil
}

func getScenarios(scanner *bufio.Scanner) ([]Scenario, error) {
	var data []Scenario
	for scanner.Scan() {
		line := scanner.Text()
		sc, err := NewScenario(line)
		if err != nil {
			return nil, err
		}
		data = append(data, *sc)
	}
	return data, nil
}

type Scenario struct {
	Bucket        int
	Map           string
	MapWidth      int
	MapHeight     int
	StartX        int
	StartY        int
	GoalX         int
	GoalY         int
	OptimalLength float64
}

func NewScenario(line string) (*Scenario, error) {
	data := strings.Fields(line)
	if len(data) != 9 {
		return nil, fmt.Errorf("invalid scenario line: '%s'", line)
	}
	s := &Scenario{}
	var err error

	s.Bucket, err = parseInt(data[0], "bucket")
	if err != nil {
		return nil, err
	}
	s.Map = data[1]
	s.MapWidth, err = parseInt(data[2], "map-width")
	if err != nil {
		return nil, err
	}
	s.MapHeight, err = parseInt(data[3], "map-width")
	if err != nil {
		return nil, err
	}
	s.StartX, err = parseInt(data[4], "start-x")
	if err != nil {
		return nil, err
	}
	s.StartY, err = parseInt(data[5], "start-y")
	if err != nil {
		return nil, err
	}
	s.GoalX, err = parseInt(data[6], "goal-x")
	if err != nil {
		return nil, err
	}
	s.GoalY, err = parseInt(data[7], "goal-y")
	if err != nil {
		return nil, err
	}
	s.OptimalLength, err = strconv.ParseFloat(data[8], 64)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func parseInt(raw string, name string) (int, error) {
	i, err := strconv.ParseInt(raw, 10, 0)
	if err != nil {
		return -1, fmt.Errorf("invalid %s: '%s'", name, raw)
	}
	return int(i), nil
}
