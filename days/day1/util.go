package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const SEPARATOR = "   "

func locationComponents(line string) (int, int, error) {
	parts := strings.Split(line, SEPARATOR)

	if len(parts) > 2 {
		return 0, 0, fmt.Errorf("Too many parts")
	}

	if len(parts) <= 0 {
		return 0, 0, fmt.Errorf("No parts")
	}

	leftID, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	rightID, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return leftID, rightID, nil
}

func lists(data io.Reader) ([]int, []int, error) {
	scanner := bufio.NewScanner(data)

	leftIDs := []int{}
	rightIDs := []int{}

	for scanner.Scan() {
		leftID, rightID, err := locationComponents(scanner.Text())
		if err != nil {
			return nil, nil, err
		}

		leftIDs = append(leftIDs, leftID)
		rightIDs = append(rightIDs, rightID)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftIDs, rightIDs, nil
}
