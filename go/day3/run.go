package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type game struct {
	opp string
	me  string
}

func main() {
	fmt.Println("Day 3!")
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	startTime := time.Now()
	reader := bufio.NewScanner(f)

	file := []game{}
	for reader.Scan() {
		line := reader.Text()
		var oppTemp string
		var meTemp string
		_, err := fmt.Sscanf(line, "%s %s", &oppTemp, &meTemp)
		if err != nil {
			fmt.Println("Error 1")
			break
		}
		file = append(file, game{opp: oppTemp, me: meTemp})
	}

	fmt.Println("Read time:", time.Since(startTime))
	// part 1
	startTime = time.Now()

	highest := 0
	for _, c := range file {
		gamePoints := 0
		if c.me == "X" {
			gamePoints = gamePoints + 1
			if c.opp == "A" {
				gamePoints = gamePoints + 3
			}
			if c.opp == "C" {
				gamePoints = gamePoints + 6
			}
		} else if c.me == "Y" {
			gamePoints = gamePoints + 2
			if c.opp == "B" {
				gamePoints = gamePoints + 3
			}
			if c.opp == "A" {
				gamePoints = gamePoints + 6
			}
		} else if c.me == "Z" {
			gamePoints = gamePoints + 3
			if c.opp == "C" {
				gamePoints = gamePoints + 3
			}
			if c.opp == "B" {
				gamePoints = gamePoints + 6
			}
		}
		highest = highest + gamePoints
	}
	fmt.Println("Part 1 answer: ", highest)

	fmt.Println("Part 1 duration:", time.Since(startTime))
	//part 2
	startTime = time.Now()

	highest2 := 0
	for _, c := range file {
		gamePoints := 0
		if c.me == "X" {
			if c.opp == "A" {
				gamePoints = gamePoints + 3
			}
			if c.opp == "B" {
				gamePoints = gamePoints + 1
			}
			if c.opp == "C" {
				gamePoints = gamePoints + 2
			}
		} else if c.me == "Y" {
			gamePoints = gamePoints + 3
			if c.opp == "A" {
				gamePoints = gamePoints + 1
			}
			if c.opp == "B" {
				gamePoints = gamePoints + 2
			}
			if c.opp == "C" {
				gamePoints = gamePoints + 3
			}
		} else if c.me == "Z" {
			gamePoints = gamePoints + 6
			if c.opp == "A" {
				gamePoints = gamePoints + 2
			}
			if c.opp == "B" {
				gamePoints = gamePoints + 3
			}
			if c.opp == "C" {
				gamePoints = gamePoints + 1
			}
		}
		highest2 = highest2 + gamePoints
	}

	fmt.Printf("Part 2 answer: %d\n", highest2)

	fmt.Println("Part 2 duration:", time.Since(startTime))
}
