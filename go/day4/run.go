package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Day 4!")
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	startTime := time.Now()
	reader := bufio.NewScanner(f)

	file := []string{}
	for reader.Scan() {
		line := reader.Text()
		file = append(file, line)
	}

	fmt.Println("Read time:", time.Since(startTime))
	// part 1
	startTime = time.Now()

	highest := 0
	for _, c := range file {
		size := len(c) / 2
		first := c[0:size]
		second := c[size:]
		for _, l := range first {
			if strings.ContainsRune(second, l) {
				highest = highest + getPriority(l)
				break
			}
		}
	}
	fmt.Println("Part 1 answer: ", highest)

	fmt.Println("Part 1 duration:", time.Since(startTime))
	//part 2
	startTime = time.Now()

	highest2 := 0
	for i := 0; i < len(file); i = i + 3 {
		first := file[i]
		second := file[i+1]
		third := file[i+2]
		for _, l := range first {
			if strings.ContainsRune(second, l) && strings.ContainsRune(third, l) {
				highest2 = highest2 + getPriority(l)
				break
			}
		}
	}

	fmt.Printf("Part 2 answer: %d\n", highest2)

	fmt.Println("Part 2 duration:", time.Since(startTime))
}

func getPriority(r rune) int {
	if r == 'a' {
		return 1
	}
	if r == 'b' {
		return 2
	}
	if r == 'c' {
		return 3
	}
	if r == 'd' {
		return 4
	}
	if r == 'e' {
		return 5
	}
	if r == 'f' {
		return 6
	}
	if r == 'g' {
		return 7
	}
	if r == 'h' {
		return 8
	}
	if r == 'i' {
		return 9
	}
	if r == 'j' {
		return 10
	}
	if r == 'k' {
		return 11
	}
	if r == 'l' {
		return 12
	}
	if r == 'm' {
		return 13
	}
	if r == 'n' {
		return 14
	}
	if r == 'o' {
		return 15
	}
	if r == 'p' {
		return 16
	}
	if r == 'q' {
		return 17
	}
	if r == 'r' {
		return 18
	}
	if r == 's' {
		return 19
	}
	if r == 't' {
		return 20
	}
	if r == 'u' {
		return 21
	}
	if r == 'v' {
		return 22
	}
	if r == 'w' {
		return 23
	}
	if r == 'x' {
		return 24
	}
	if r == 'y' {
		return 25
	}
	if r == 'z' {
		return 26
	}
	if r == 'A' {
		return 27
	}
	if r == 'B' {
		return 28
	}
	if r == 'C' {
		return 29
	}
	if r == 'D' {
		return 30
	}
	if r == 'E' {
		return 31
	}
	if r == 'F' {
		return 32
	}
	if r == 'G' {
		return 33
	}
	if r == 'H' {
		return 34
	}
	if r == 'I' {
		return 35
	}
	if r == 'J' {
		return 36
	}
	if r == 'K' {
		return 37
	}
	if r == 'L' {
		return 38
	}
	if r == 'M' {
		return 39
	}
	if r == 'N' {
		return 40
	}
	if r == 'O' {
		return 41
	}
	if r == 'P' {
		return 42
	}
	if r == 'Q' {
		return 43
	}
	if r == 'R' {
		return 44
	}
	if r == 'S' {
		return 45
	}
	if r == 'T' {
		return 46
	}
	if r == 'U' {
		return 47
	}
	if r == 'V' {
		return 48
	}
	if r == 'W' {
		return 49
	}
	if r == 'X' {
		return 50
	}
	if r == 'Y' {
		return 51
	}
	if r == 'Z' {
		return 52
	}
	return 1000
}
