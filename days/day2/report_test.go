package day2_test

import (
	"testing"

	"github.com/mattismoel/aoc25/days/day2"
)

func TestIsAscending(t *testing.T) {
	type test struct {
		name          string
		r             day2.Report
		wantAscending bool
	}

	tests := []test{
		{
			name:          "Simple",
			r:             []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantAscending: true,
		},
		{
			name:          "Descending",
			r:             []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			wantAscending: false,
		},
		{
			name:          "One off",
			r:             []int{1, 2, 3, 4, 5, 4, 7, 8, 9, 10},
			wantAscending: false,
		},
		{
			name:          "Duplicate entries",
			r:             []int{1, 2, 3, 3, 5, 6, 7, 7, 9, 10},
			wantAscending: true,
		},
		{
			name:          "Zero entries",
			r:             []int{},
			wantAscending: true,
		},
		{
			name:          "One entry",
			r:             []int{1},
			wantAscending: true,
		},
		{
			name:          "Negative entries ascending",
			r:             []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
			wantAscending: true,
		},
		{
			name:          "Negative entries ascending",
			r:             []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			wantAscending: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAscending := tt.r.IsAscending()

			if gotAscending != tt.wantAscending {
				t.Fatalf("got %v, wantAscending %v\n", gotAscending, tt.wantAscending)
			}
		})
	}
}

func TestDescending(t *testing.T) {
	type test struct {
		name     string
		r        day2.Report
		wantDesc bool
	}

	tests := []test{
		{
			name:     "Simple",
			r:        []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			wantDesc: true,
		},
		{
			name:     "Ascending",
			r:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantDesc: false,
		},
		{
			name:     "One off",
			r:        []int{10, 9, 8, 18, 6, 5, 4, 3, 2, 1},
			wantDesc: false,
		},
		{
			name:     "Duplicate entries",
			r:        []int{10, 9, 8, 7, 7, 5, 4, 3, 3, 1},
			wantDesc: true,
		},
		{
			name:     "Zero entries",
			r:        []int{},
			wantDesc: true,
		},
		{
			name:     "One entry",
			r:        []int{1},
			wantDesc: true,
		},
		{
			name:     "Negative entries descending",
			r:        []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			wantDesc: true,
		},
		{
			name:     "Negative entries ascending",
			r:        []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
			wantDesc: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDesc := tt.r.IsDescending()
			if gotDesc != tt.wantDesc {
				t.Fatalf("got %v, want %v\n", gotDesc, tt.wantDesc)
			}
		})
	}
}
