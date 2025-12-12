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
			} else {

				l, _ := strconv.Atoi(s[:len(s)/2])
				r, _ := strconv.Atoi(s[len(s)/2:])

				if l == r {
					total += i
				}
			}
		}

	}

	fmt.Printf("Total of invalid Ids: %d\n", total)
}

func readFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return string(content)
}
