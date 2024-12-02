package day2

import (
	"fmt"
	"io"
)

type Part1 struct{}

func (p Part1) Do(data io.Reader) error {
	reports, err := reportsFromData(data)
	if err != nil {
		return err
	}

	safe, _ := categorizeReports(reports)

	fmt.Printf("Safe reports: %d/%d\n", len(safe), len(reports))

	return nil
}
