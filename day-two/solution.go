package daytwo

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne() {
	input := readFile("day-two/input.txt")
	lines := strings.Split(input, ",")
	total := 0

	for _, line := range lines {
		vals := strings.Split(line, "-")
		left, _ := strconv.Atoi(vals[0])
		right, _ := strconv.Atoi(vals[1])

		for i := left; i <= right; i++ {
			s := strconv.Itoa(i)

			if len(s)%2 != 0 {
				continue
			}
			l, _ := strconv.Atoi(s[:len(s)/2])
			r, _ := strconv.Atoi(s[len(s)/2:])

			if l == r {
				total += i
			}
		}

	}

	fmt.Printf("Total of invalid Ids: %d\n", total)
}
func PartTwo() {
	// read input.txt
	input := readFile("day-two/input.txt")
	lines := strings.Split(input, ",")
	total := 0

	for _, line := range lines {
		vals := strings.Split(line, "-")
		left, _ := strconv.Atoi(vals[0])
		right, _ := strconv.Atoi(vals[1])

		for i := left; i <= right; i++ {
			s := strconv.Itoa(i)

			if len(s)%2 != 0 {
				//check if the fist number is repeated throughout the number
				first := string(s[0])
				found := true
				for j := 1; j < len(s); j++ {
					if string(s[j]) != first {
						found = false
						break
					}
				}

				if found {
					fmt.Println(i)
					total += i
				}
				continue
			}
			l, _ := strconv.Atoi(s[:len(s)/2])
			r, _ := strconv.Atoi(s[len(s)/2:])

			if l == r {
				total += i
			}
		}
	}
	fmt.Printf("Part Two Total of invalid Ids: %d\n", total)
}

func readFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return string(content)
}
