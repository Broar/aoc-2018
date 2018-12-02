package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	var result int64
	frequencies := strings.Split(string(raw), "\n")

	for _, frequency := range frequencies {
		value, err := strconv.ParseInt(frequency, 10, 64)
		if err != nil {
			log.Fatalf("unable to convert frequency from string to integer: %v", err)
		}

		result += value
	}

	fmt.Println(result)
}
