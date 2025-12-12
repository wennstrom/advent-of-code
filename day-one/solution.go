package dayone

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne() {
	input := readFile("day-one/input.txt")
	clicks := 0
	pos := 50

	for _, l := range input {
		//get first character in line
		dir := string(l[0])
		//distance
		dis, _ := strconv.Atoi(strings.Split(l, dir)[1])
		//get position after rotation
		pos = rotate(dir, pos, dis)
		if pos == 0 {
			clicks += 1
		}
	}
	fmt.Printf("Part One Password:%d\n", clicks)
}

func PartTwo() {
	input := readFile("day-one/input.txt")
	totalClicks := 0
	pos := 50

	for _, l := range input {
		//get first character in line
		dir := string(l[0])
		//distance
		dis, _ := strconv.Atoi(strings.Split(l, dir)[1])
		//get position after rotation
		newPos, clicks := rotateWithClick(dir, pos, dis)
		pos = newPos
		totalClicks += clicks
	}
	fmt.Printf("Part Two Password:%d\n", totalClicks)
}

func rotate(dir string, pos, dis int) int {
	if dir == "R" {
		return (pos + dis) % 100
	}

	return (pos - dis + 100) % 100
}

func rotateWithClick(dir string, pos, dis int) (int, int) {

	r := dis % 100

	if dir == "R" {

		p := (pos + dis) % 100
		c := (pos + dis) / 100

		if (pos + r) > 100 {
			c++
		}

		return p, c
	}

	if (pos - dis + 100) >= 0 {
		return pos - dis, 0
	}

	p := (pos - dis) % 100
	c := dis / 100

	if pos-r < 0 {
		c++
	}
	return p, c
}

func readFile(filename string) []string {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return strings.Split(string(content), "\n")

}
