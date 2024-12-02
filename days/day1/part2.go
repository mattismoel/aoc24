package day1

import (
	"fmt"
	"io"
	"log"

	"github.com/mattismoel/aoc25/internal/number"
)

type Part2 struct{}

func (p Part2) Do(data io.Reader) error {

	leftIDs, rightIDs, err := lists(data)
	if err != nil {
		return err
	}

	similarityScores, err := similarityScores(leftIDs, rightIDs)
	if err != nil {
		log.Fatal(err)
	}

	sum := number.Sum(similarityScores)

	fmt.Printf("Similarity Score Sum: %d\n", sum)

	return nil
}

func similarityScores(leftIDs []int, rightIDs []int) ([]int, error) {
	if len(leftIDs) != len(rightIDs) {
		return nil, fmt.Errorf("Lists are not the same length")
	}

	similarityScores := []int{}

	for _, leftID := range leftIDs {
		existCount := 0
		for _, rightID := range rightIDs {
			if rightID == leftID {
				existCount++
			}
		}

		similarityScores = append(similarityScores, leftID*existCount)
	}

	return similarityScores, nil
}
