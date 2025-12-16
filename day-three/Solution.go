package daythree

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne() {
	banks := getInput("day-three/input.txt")
	total := 0
	for _, bank := range banks {
		jolt := 0
		cur := 0
		for _, digit := range bank {
			i := int(digit - '0')
			if cur == 0 {
				cur = i
				continue
			}
			val, _ := strconv.Atoi(strconv.Itoa(cur) + strconv.Itoa(i))

			if val > jolt {
				jolt = val
			}
			if i > cur {
				cur = i
			}
		}
		total += jolt
	}
	fmt.Printf("Total: %d\n", total)
}
func PartTwo() {
	banks := getInput("day-three/input.txt")
	total := 0
	for _, bank := range banks {
		jolt := 0
		cur := 0
		x := 0
		for i := 0; i < len(bank); i++ {
			digit := int(bank[i] - '0')
			if len(bank)-i < 12 {
				break
			}

			if x == 0 {
				x = digit
				cur, _ = strconv.Atoi(bank[i : i+12])
				continue
			}
			if digit > x {
				x = digit
				cur, _ = strconv.Atoi(bank[i : i+12])
			}
		}
		total += jolt
	}

}
func getInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return strings.Split(string(content), "\n")
}
