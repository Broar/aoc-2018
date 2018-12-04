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

	var ids []string
	for _, line := range lines {
		counts := make(map[rune]int)

		for _, letter := range line {
			if _, exists := counts[letter]; exists {
				counts[letter]++
			} else {
				counts[letter] = 1
			}
		}

		twos := 0
		threes := 0
		for _, count := range counts {
			if count == 2 {
				twos++
			} else if count == 3 {
				threes++
			}
		}

		if (twos == 1 && threes == 1) || (twos == 0 && threes == 1) || (twos == 1 && threes == 0) {
			ids = append(ids, line)
		}
	}

	fmt.Println(ids)

	for _, id1 := range lines {
		for _, id2 := range lines {
			candidate := strings.Builder{}
			difference := 0

			for i := 0; i < len(id1); i++ {
				if difference > 1 {
					break
				}

				if id1[i] == id2[i] {
					if err := candidate.WriteByte(id1[i]); err != nil {
						log.Fatalf("unable to append byte to string: %v", err)
					}

				} else {
					difference++
				}
			}

			if difference == 1 {
				fmt.Println(candidate.String())
				return
			}
		}
	}
}
