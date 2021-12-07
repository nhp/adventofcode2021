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

func main() {
	fmt.Println("First of december!")
	depths, err := readDepths("1.data")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(depths))
	var last = 0
	var count = 0
	for _, s := range depths {
		if last == 0 {
			last = s
			fmt.Println(last)
			continue
		}
		if last < s {
			fmt.Println(last, s, count)
			count++
		}
		last = s
		fmt.Println(count)
	}
	fmt.Println(count)
}
