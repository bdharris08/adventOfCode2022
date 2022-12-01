package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	calories int
}

type byMostCalories []Elf

func (s byMostCalories) Len() int {
	return len(s)
}
func (s byMostCalories) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byMostCalories) Less(i, j int) bool {
	return s[i].calories > s[j].calories
}

// Part 1
func topElfCalories(elfs []Elf) int {
	// assumes sorted
	return elfs[0].calories
}

// Part 2
func top3ElfCalories(elfs []Elf) int {
	// assumes sorted
	sum := 0
	for i := 0; i < 3; i++ {
		sum += elfs[i].calories
	}
	return sum
}

func main() {
	fmt.Println("day 1")
	elfs := []Elf{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elf := Elf{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			elfs = append(elfs, elf)
			elf = Elf{}
		}

		calories, err := strconv.Atoi(scanner.Text())
		if err == nil {
			elf.calories += calories
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(byMostCalories(elfs))
	fmt.Println("top elf: ", topElfCalories(elfs))
	fmt.Println("top 3 sum: ", top3ElfCalories(elfs))
}
