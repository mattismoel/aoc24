package manager

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/mattismoel/aoc25/days/day1"
	"github.com/mattismoel/aoc25/days/day2"
)

type puzzler interface {
	Do(data io.Reader) error
}

type manager struct {
	dir     string
	puzzles map[int][]puzzler
}

func NewManager(dir string) *manager {
	dir = path.Dir(dir)

	m := &manager{
		dir: dir,
		puzzles: map[int][]puzzler{
			1: {
				day1.Part1{},
				day1.Part2{},
			},
			2: {
				day2.Part1{},
				day2.Part2{},
			},
		},
	}

	return m
}

func (m manager) dayData(day int) (io.ReadCloser, error) {
	filePath := fmt.Sprintf("%s/day%d.txt", m.dir, day)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (m manager) DoPuzzle(day int) error {
	data, err := m.dayData(day)
	if err != nil {
		return err
	}

	puzzles, ok := m.puzzles[day]
	if !ok {
		return fmt.Errorf("Puzzle for day %d does not exist", day)
	}

	fmt.Printf("\nDAY %d:\n", day)

	var buffer bytes.Buffer
	tee := io.TeeReader(data, &buffer)

	_, err = io.ReadAll(tee)
	if err != nil {
		return err
	}

	for i, p := range puzzles {
		reader := bytes.NewReader(buffer.Bytes())
		fmt.Printf("Part %d/%d:\n", i+1, len(puzzles))

		err = p.Do(reader)
		if err != nil {
			return fmt.Errorf("Could not do puzzle: %v", err)
		}

		fmt.Println()
	}

	return nil
}
