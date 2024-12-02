package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Report []int

func reportsFromData(data io.Reader) ([]Report, error) {
	reports := []Report{}
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		r, err := reportFromLine(scanner.Text())
		if err != nil {
			return nil, err
		}

		reports = append(reports, r)
	}

	return reports, nil
}

func reportFromLine(line string) (Report, error) {
	ints := []int{}
	parts := strings.Split(line, " ")

	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}

	return ints, nil

}

func categorizeReports(reports []Report) ([]Report, []Report) {
	safeReports := []Report{}
	unsafeReports := []Report{}

	for _, r := range reports {
		if r.IsSafe() {
			safeReports = append(safeReports, r)
			continue
		}

		unsafeReports = append(unsafeReports, r)
	}

	return safeReports, unsafeReports
}

func (r Report) IsSafe() bool {
	if (!r.IsDescending()) && (!r.IsAscending()) {
		return false
	}

	for i, num := range r {
		if i == 0 {
			continue
		}

		prevNum := r[i-1]

		if !isSafeDiff(prevNum, num) {
			return false
		}
	}

	return true
}

func (r Report) IsAscending() bool {
	for i, num := range r {
		if i == 0 {
			continue
		}

		prevNum := r[i-1]

		if !isAscending(prevNum, num) {
			return false
		}
	}

	return true
}

func (r Report) IsDescending() bool {
	for i, num := range r {
		if i == 0 {
			continue
		}

		prevNum := r[i-1]

		if !isDescending(prevNum, num) {
			return false
		}
	}

	return true
}
