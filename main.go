package main

import (
	"flag"
	"log"

	"github.com/mattismoel/aoc25/internal/manager"
)

const DATA_DIR = "./data/"

func main() {
	day := flag.Int("day", 1, "The day of Advent Of Code puzzle to do")

	flag.Parse()

	manager := manager.NewManager(DATA_DIR)

	err := manager.DoPuzzle(*day)
	if err != nil {
		log.Fatal(err)
	}
}
