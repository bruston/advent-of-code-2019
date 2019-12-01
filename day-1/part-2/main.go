package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	input := flag.String("f", "input.txt", "location of input.txt file")
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		log.Fatalf("Unable to open input file: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		n := 0
		if _, err := fmt.Sscan(s.Text(), &n); err != nil {
			log.Fatalf("error while scanning input: %v", err)
		}
		fuel := n/3 - 2
		sum += fuel
		for {
			fuel = fuel/3 - 2
			if fuel <= 0 {
				break
			}
			sum += fuel
		}
	}
	if s.Err() != nil {
		log.Fatalf("Error while reading from input file: %v", err)
	}
	fmt.Println(sum)
}
