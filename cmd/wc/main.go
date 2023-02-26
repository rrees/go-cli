package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	counter := 0

	scanner := bufio.NewScanner(r)

	if countLines == false {
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		counter += 1
	}

	return counter
}

func main() {
	lineFlag := flag.Bool("l", false, "Count lines")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lineFlag))
}
