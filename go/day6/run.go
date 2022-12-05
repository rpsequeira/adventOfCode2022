package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type move struct {
	n_crates int
	orig     int
	dest     int
}

func main() {
	fmt.Println("Day 6!")
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	startTime := time.Now()
	reader := bufio.NewScanner(f)

	state := [][]rune{
		// 0
		{'F', 'T', 'C', 'L', 'R', 'P', 'G', 'Q'},
		// 1
		{'N', 'Q', 'H', 'W', 'R', 'F', 'S', 'J'},
		// 2
		{'F', 'B', 'H', 'W', 'P', 'M', 'Q'},
		// 3
		{'V', 'S', 'T', 'D', 'F'},
		// 4
		{'Q', 'L', 'D', 'W', 'V', 'F', 'Z'},
		// 5
		{'Z', 'C', 'L', 'S'},
		// 6
		{'Z', 'B', 'M', 'V', 'D', 'F'},
		// 7
		{'T', 'J', 'B'},
		// 8
		{'Q', 'N', 'B', 'G', 'L', 'S', 'P', 'H'},
	}
	file := []move{}
	for reader.Scan() {
		line := reader.Text()
		var a int
		var b int
		var c int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &a, &b, &c)
		if err != nil {
			fmt.Println("Error 1")
			break
		}
		file = append(file, move{n_crates: a, orig: b - 1, dest: c - 1})
	}

	fmt.Println("Read time:", time.Since(startTime))
	// part 1
	startTime = time.Now()
	for _, m := range file {
		index := len(state[m.orig]) - m.n_crates
		t := state[m.orig][index:]
		state[m.orig] = state[m.orig][0:index]
		for i := len(t) - 1; i >= 0; i-- {
			state[m.dest] = append(state[m.dest], t[i])
		}
	}
	fmt.Println("Part 1 answer: ")
	for _, s := range state {
		fmt.Print(string(s[len(s)-1]))
	}
	fmt.Println()
	fmt.Println("Part 1 duration:", time.Since(startTime))
	//part 2
	startTime = time.Now()
	state = [][]rune{
		// 0
		{'F', 'T', 'C', 'L', 'R', 'P', 'G', 'Q'},
		// 1
		{'N', 'Q', 'H', 'W', 'R', 'F', 'S', 'J'},
		// 2
		{'F', 'B', 'H', 'W', 'P', 'M', 'Q'},
		// 3
		{'V', 'S', 'T', 'D', 'F'},
		// 4
		{'Q', 'L', 'D', 'W', 'V', 'F', 'Z'},
		// 5
		{'Z', 'C', 'L', 'S'},
		// 6
		{'Z', 'B', 'M', 'V', 'D', 'F'},
		// 7
		{'T', 'J', 'B'},
		// 8
		{'Q', 'N', 'B', 'G', 'L', 'S', 'P', 'H'},
	}

	for _, m := range file {
		index := len(state[m.orig]) - m.n_crates
		t := state[m.orig][index:]
		state[m.orig] = state[m.orig][0:index]
		state[m.dest] = append(state[m.dest], t...)
	}
	fmt.Println("Part 1 answer: ")
	for _, s := range state {
		fmt.Print(string(s[len(s)-1]))
	}
	fmt.Println()
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
