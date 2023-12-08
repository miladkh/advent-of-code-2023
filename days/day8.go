package days

import (
	"bufio"
	"fmt"
	"os"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func FindDistance(start string, end func(string) bool, points map[string][]string, inst string) int {
	current := points[start]
	key := start
	endloop := false
	sum := 0
	for {
		for _, v := range inst {
			if end(key) {
				endloop = true
				break
			}
			if v == 76 {
				key = current[0]
				current = points[current[0]]
			} else {
				key = current[1]
				current = points[current[1]]
			}
			sum++
		}
		if endloop {
			break
		}
	}
	return sum
}

func Day8Part1() {
	f, err := os.Open("data/day8.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	inst := scanner.Text()
	scanner.Scan()
	points := map[string][]string{}
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			points[s[:3]] = []string{s[7:10], s[12:15]}
		}
	}
	fmt.Println(FindDistance("AAA", func(s string) bool { return s == "ZZZ" }, points, inst))
}

func Day8Part2() {
	f, err := os.Open("data/day8.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	inst := scanner.Text()
	scanner.Scan()
	points := map[string][]string{}
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			points[s[:3]] = []string{s[7:10], s[12:15]}
		}
	}
	nums := []int{}
	for k, _ := range points {
		if k[2] == 65 {
			nums = append(nums, FindDistance(k, func(s string) bool { return s[2] == 90 }, points, inst))
		}
	}
	fmt.Println(LCM(nums[0], nums[1], nums[2:]...))
}
