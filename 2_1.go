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
	var forward, depth = 0, 0
	for scanner.Scan() {
		x := strings.Fields(scanner.Text())
		for _, value := range x {
			var c = 0
			switch value {
			case "forward":
				c, _ = strconv.Atoi(x[1])
				forward += c

			case "down":
				c, _ = strconv.Atoi(x[1])
				depth += c
			case "up":
				c, _ = strconv.Atoi(x[1])
				depth -= c
			}
		}
	}
	return depth, forward, scanner.Err()
}

func main() {
	depth, forward, _ := readDirection("2_1.data")
	fmt.Printf("depth: %d\n", depth)
	fmt.Printf("forward: %d\n", forward)
	fmt.Printf("Solution: %d\n", forward*depth)
}
