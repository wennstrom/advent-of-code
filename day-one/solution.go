package dayone

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution() {
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
		fmt.Println("pos:", pos, "dir:", dir, "dis:", dis, "clicks:", clicks)
		if pos == 0 {
			clicks += 1
		}
	}
	fmt.Printf("Password:%d\n", clicks)
}

func rotate(dir string, pos, dis int) int {
	if dir == "R" {
		return (pos + dis) % 100
	}

	return (pos - dis + 100) % 100
}

func readFile(filename string) []string {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return strings.Split(string(content), "\n")

}
