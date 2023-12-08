package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Day1Part1() {
	f, err := os.Open("data/day1.txt")
	check(err)
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		first := -1
		last := -1
		for i := 0; i < len(s); i++ {
			if s[i] >= 48 && s[i] <= 57 {
				if first == -1 {
					first = int(s[i]) - 48
				}
				last = int(s[i]) - 48
			}
		}
		sum += last + (10 * first)
	}
	fmt.Println(sum)
	f.Close()
}

func digitize(digit string) int {
	r := -1
	if strings.HasSuffix(digit, "one") {
		r = 1
	} else if strings.HasSuffix(digit, "two") {
		r = 2
	} else if strings.HasSuffix(digit, "three") {
		r = 3
	} else if strings.HasSuffix(digit, "four") {
		r = 4
	} else if strings.HasSuffix(digit, "five") {
		r = 5
	} else if strings.HasSuffix(digit, "six") {
		r = 6
	} else if strings.HasSuffix(digit, "seven") {
		r = 7
	} else if strings.HasSuffix(digit, "eight") {
		r = 8
	} else if strings.HasSuffix(digit, "nine") {
		r = 9
	}
	return r
}

func Day1Part2() {
	f, err := os.Open("data/day1.txt")
	check(err)
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		first := -1
		last := -1
		digit := ""
		for i := 0; i < len(s); i++ {
			if s[i] >= 48 && s[i] <= 57 {
				if first == -1 {
					first = int(s[i]) - 48
				}
				last = int(s[i]) - 48
				digit = ""
			} else {
				digit += string(s[i])
			}
			d := digitize(digit)
			if d >= 0 {
				last = d
				if first == -1 {
					first = d
				}
			}
		}
		sum += last + (10 * first)
	}
	fmt.Println(sum)
	f.Close()
}
