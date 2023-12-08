package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day4Part1() {
	f, err := os.Open("data/day4.txt")
	check(err)
	sum := 0.0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		s1 := strings.Split(s, ":")[1]
		parts := strings.Split(s1, "|")
		winning := strings.Split(parts[0], " ")
		nums := strings.Split(parts[1], " ")
		mwinning := make(map[string]int)
		for w := range winning {
			if len(strings.Trim(winning[w], " ")) > 0 {
				mwinning[winning[w]] = 1
			}
		}
		counter := 0
		for n := range nums {
			if mwinning[nums[n]] == 1 {
				counter++
			}
		}
		fmt.Println(s)
		if counter > 0 {
			fmt.Println(math.Pow(2, float64(counter-1)))
			sum += math.Pow(2, float64(counter-1))
		}
	}
	fmt.Println(sum)
}

type day4Line struct {
	card    int
	counter int
}

func Day4Part2() {
	f, err := os.Open("data/day4.txt")
	check(err)
	sum := 0
	scores := make(map[int]int)
	var cards []day4Line
	scanner := bufio.NewScanner(f)
	pattern := "\\s+"
	regex := regexp.MustCompile(pattern)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		ss := strings.Split(s, ":")
		s1 := ss[1]
		card, err := strconv.Atoi(regex.Split(ss[0], -1)[1])
		check(err)
		parts := strings.Split(s1, "|")
		winning := strings.Split(parts[0], " ")
		nums := strings.Split(strings.Trim(parts[1], " "), " ")
		mwinning := make(map[string]int)
		for w := range winning {
			if len(strings.Trim(winning[w], " ")) > 0 {
				mwinning[winning[w]] = 1
			}
		}
		counter := 0
		for n := range nums {
			if mwinning[nums[n]] == 1 {
				counter++
			}
		}
		cards = append(cards, day4Line{card: card, counter: counter})
		scores[card] = 1
	}
	for l := range cards {
		c := cards[l]
		for i := 1; i <= c.counter; i++ {
			scores[c.card+i] += scores[c.card]
		}
	}
	for sss := range scores {
		sum += scores[sss]
	}
	fmt.Println(sum)
}
