package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readDirection(path string) (int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var horizontal, depth, aim = 0, 0, 0
	for scanner.Scan() {
		x := strings.Fields(scanner.Text())
		for _, value := range x {
			var c = 0
			switch value {
			case "forward":
				c, _ = strconv.Atoi(x[1])
				depth += aim * c
				horizontal += c
			case "down":
				c, _ = strconv.Atoi(x[1])
				aim += c
			case "up":
				c, _ = strconv.Atoi(x[1])
				aim -= c
			}
		}
	}
	return depth, horizontal, scanner.Err()
}

func main() {
	depth, horizontal, _ := readDirection("2_1.data")
	fmt.Printf("depth: %d\n", depth)
	fmt.Printf("horizontal: %d\n", horizontal)
	fmt.Printf("Solution: %d\n", horizontal*depth)
}
