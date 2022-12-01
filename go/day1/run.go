package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Day 1!")
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	startTime := time.Now()
	reader := bufio.NewScanner(f)

	file := []int{}
	caloriesPerElf := 0
	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			file = append(file, caloriesPerElf)
			caloriesPerElf = 0
		} else {
			var calories string
			_, err := fmt.Sscanf(line, "%s", &calories)
			if err != nil {
				fmt.Println("Error 1")
				break
			}
			c, err := strconv.Atoi(calories)
			if err != nil {
				fmt.Println("Error 2")
				break
			}
			caloriesPerElf = caloriesPerElf + c
		}

	}

	fmt.Println("Read time:", time.Since(startTime))
	startTime = time.Now()
	// part 1

	highest := 0
	for _, c := range file {
		if highest < c {
			highest = c
		}
	}
	fmt.Println("Part 1 answer: ", highest)

	fmt.Println("Part 1 duration:", time.Since(startTime))
	startTime = time.Now()
	//part 2

	highest2 := 0

	sort.Ints(file)
	for i := len(file) - 1; i > len(file)-4; i-- {
		highest2 = highest2 + file[i]
	}

	fmt.Printf("Part 2 answer: %d\n", highest2)

	fmt.Println("Part 2 duration:", time.Since(startTime))
}
