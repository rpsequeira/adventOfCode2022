package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	startTime := time.Now()
	reader := bufio.NewScanner(f)

	file := []int{}
	for reader.Scan() {
		var number int
		var direction string
		_, err := fmt.Sscanf(reader.Text(), "%s %d", &direction, &number)
		if err != nil {
			break
		}

		file = append(file, number)
	}

	fmt.Println("Read time:", time.Since(startTime))
	startTime = time.Now()
	// part 1

	previous := -1
	counter := 0
	for i := range file {
		if previous != -1 && previous < file[i] {
			counter++
		}

		previous = file[i]
	}
	fmt.Println(counter)

	fmt.Println("Part 1:", time.Since(startTime))
	startTime = time.Now()
	//part 2

	previous = -1
	counter = 0
	for i := 0; i < len(file)-2; i++ {
		sum := file[i] + file[i+1] + file[i+2]
		if previous != -1 && previous < sum {
			counter++
		}

		previous = sum
	}
	fmt.Println(counter)

	fmt.Println("Part 2:", time.Since(startTime))
}
