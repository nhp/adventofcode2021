package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readDepths(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var depths []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, x)
	}
	return depths, scanner.Err()
}

func getSlidingWindowSum(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("First of december. Second puzzle!")
	depths, err := readDepths("1_2.data")
	if err != nil {
		log.Fatal(err)
	}
	var last = 0
	var count = 0
	for i, _ := range depths {
		if i+2 >= len(depths) {
			continue
		}
		var sum = getSlidingWindowSum(depths[i], depths[i+1], depths[i+2])
		if last == 0 {
			last = sum
			continue
		}
		if last < sum {
			count++
		}
		last = sum
		fmt.Println(sum)
	}
	fmt.Println(count)
}
