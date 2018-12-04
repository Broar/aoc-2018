package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	lines := strings.Split(string(raw), "\n")

	twos := 0
	threes := 0

	for _, line := range lines {
		counts := make(map[rune]int)

		for _, letter := range line {
			if _, exists := counts[letter]; exists {
				counts[letter]++
			} else {
				counts[letter] = 1
			}
		}

		seenTwo := false
		seenThree := false
		for _, count := range counts {
			if seenTwo && seenThree {
				break
			}

			if count == 2 && !seenTwo {
				twos++
				seenTwo = true
			} else if count == 3 && !seenThree {
				threes++
				seenThree = true
			}
		}
	}

	fmt.Println(twos * threes)
}
