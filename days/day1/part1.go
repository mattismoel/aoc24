package day1

import (
	"fmt"
	"io"
	"math"
	"slices"

	"github.com/mattismoel/aoc25/internal/number"
)

type Part1 struct{}

func (p Part1) Do(data io.Reader) error {
	leftIDs, rightIDs, err := lists(data)
	if err != nil {
		return err
	}

	distances, err := distances(leftIDs, rightIDs)
	if err != nil {
		return err
	}

	sum := number.Sum(distances)
	fmt.Printf("Total sum: %d\n", sum)

	return nil
}

func calculateDistanceSum(leftIDs, rightIDs []int) (int, error) {
	slices.Sort(leftIDs)
	slices.Sort(rightIDs)

	distances, err := distances(leftIDs, rightIDs)
	if err != nil {
		return 0, err
	}

	return number.Sum(distances), nil
}

func distances(leftIDs []int, rightIDs []int) ([]int, error) {
	if len(leftIDs) != len(rightIDs) {
		return nil, fmt.Errorf("Lists are not same length")
	}

	length := len(leftIDs)

	distances := []int{}

	for i := range length {
		leftID := leftIDs[i]
		rightID := rightIDs[i]

		d := math.Abs(float64(leftID) - float64(rightID))
		distances = append(distances, int(d))
	}

	return distances, nil
}
